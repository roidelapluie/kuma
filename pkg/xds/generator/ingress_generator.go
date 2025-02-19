package generator

import (
	"sort"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	model "github.com/kumahq/kuma/pkg/core/xds"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
	envoy_clusters "github.com/kumahq/kuma/pkg/xds/envoy/clusters"
	envoy_endpoints "github.com/kumahq/kuma/pkg/xds/envoy/endpoints"
	envoy_listeners "github.com/kumahq/kuma/pkg/xds/envoy/listeners"
	envoy_names "github.com/kumahq/kuma/pkg/xds/envoy/names"
	"github.com/kumahq/kuma/pkg/xds/envoy/tls"
)

const (
	IngressProxy = "ingress-proxy"

	// OriginIngress is a marker to indicate by which ProxyGenerator resources were generated.
	OriginIngress = "ingress"
)

type IngressGenerator struct {
}

func (i IngressGenerator) Generate(ctx xds_context.Context, proxy *model.Proxy) (*model.ResourceSet, error) {
	resources := model.NewResourceSet()

	destinationsPerService := i.destinations(proxy.Routing.TrafficRouteList)

	listener, err := i.generateLDS(proxy, proxy.ZoneIngress, destinationsPerService, proxy.APIVersion)
	if err != nil {
		return nil, err
	}
	resources.Add(&model.Resource{
		Name:     listener.GetName(),
		Origin:   OriginIngress,
		Resource: listener,
	})

	services := i.services(proxy)

	cdsResources, err := i.generateCDS(services, destinationsPerService, proxy.APIVersion)
	if err != nil {
		return nil, err
	}
	resources.Add(cdsResources...)

	edsResources, err := i.generateEDS(proxy, services, proxy.APIVersion)
	if err != nil {
		return nil, err
	}
	resources.Add(edsResources...)

	return resources, nil
}

// generateLDS generates one Ingress Listener
// Ingress Listener assumes that mTLS is on. Using TLSInspector we sniff SNI value.
// SNI value has service name and tag values specified with the following format: "backend{cluster=2,version=1}"
// We take all possible destinations from TrafficRoutes and generate FilterChainsMatcher for each unique destination.
// This approach has a limitation: additional tags on outbound in Universal mode won't work across different zones.
// Traffic is NOT decrypted here, therefore we don't need certificates and mTLS settings
func (i IngressGenerator) generateLDS(
	proxy *model.Proxy,
	ingress *core_mesh.ZoneIngressResource,
	destinationsPerService map[string][]envoy_common.Tags,
	apiVersion envoy_common.APIVersion,
) (envoy_common.NamedResource, error) {
	inboundListenerName := envoy_names.GetInboundListenerName(proxy.ZoneIngress.Spec.GetNetworking().GetAddress(), proxy.ZoneIngress.Spec.GetNetworking().GetPort())
	inboundListenerBuilder := envoy_listeners.NewListenerBuilder(apiVersion).
		Configure(envoy_listeners.InboundListener(inboundListenerName, ingress.Spec.GetNetworking().GetAddress(), ingress.Spec.GetNetworking().GetPort(), model.SocketAddressProtocolTCP)).
		Configure(envoy_listeners.TLSInspector())

	if len(proxy.ZoneIngress.Spec.AvailableServices) == 0 {
		inboundListenerBuilder = inboundListenerBuilder.
			Configure(envoy_listeners.FilterChain(envoy_listeners.NewFilterChainBuilder(apiVersion)))
	}

	sniUsed := map[string]bool{}

	for _, inbound := range proxy.ZoneIngress.Spec.GetAvailableServices() {
		service := inbound.Tags[mesh_proto.ServiceTag]
		destinations := destinationsPerService[service]
		destinations = append(destinations, destinationsPerService[mesh_proto.MatchAllTag]...)

		for _, destination := range destinations {
			meshDestination := destination.
				WithTags(mesh_proto.ServiceTag, service).
				WithTags("mesh", inbound.GetMesh())
			sni := tls.SNIFromTags(meshDestination)
			if sniUsed[sni] {
				continue
			}
			sniUsed[sni] = true
			inboundListenerBuilder = inboundListenerBuilder.
				Configure(envoy_listeners.FilterChain(envoy_listeners.NewFilterChainBuilder(apiVersion).
					Configure(envoy_listeners.FilterChainMatch("tls", sni)).
					Configure(envoy_listeners.TcpProxyWithMetadata(service, envoy_common.NewCluster(
						envoy_common.WithService(service),
						envoy_common.WithTags(meshDestination.WithoutTags(mesh_proto.ServiceTag)),
					))),
				))
		}
	}

	return inboundListenerBuilder.Build()
}

func (_ IngressGenerator) destinations(trs *core_mesh.TrafficRouteResourceList) map[string][]envoy_common.Tags {
	destinations := map[string][]envoy_common.Tags{}
	for _, tr := range trs.Items {
		for _, split := range tr.Spec.Conf.GetSplitWithDestination() {
			service := split.Destination[mesh_proto.ServiceTag]
			destinations[service] = append(destinations[service], split.Destination)
		}
		for _, http := range tr.Spec.Conf.Http {
			for _, split := range http.GetSplitWithDestination() {
				service := split.Destination[mesh_proto.ServiceTag]
				destinations[service] = append(destinations[service], split.Destination)
			}
		}
	}
	return destinations
}

func (_ IngressGenerator) services(proxy *model.Proxy) []string {
	var services []string
	for service := range proxy.Routing.OutboundTargets {
		services = append(services, service)
	}
	sort.Strings(services)
	return services
}

func (i IngressGenerator) generateCDS(
	services []string,
	destinationsPerService map[string][]envoy_common.Tags,
	apiVersion envoy_common.APIVersion,
) (resources []*model.Resource, _ error) {
	for _, service := range services {
		tagSlice := envoy_common.TagsSlice(append(destinationsPerService[service], destinationsPerService[mesh_proto.MatchAllTag]...))
		tagKeySlice := tagSlice.ToTagKeysSlice().Transform(envoy_common.Without(mesh_proto.ServiceTag), envoy_common.With("mesh"))
		edsCluster, err := envoy_clusters.NewClusterBuilder(apiVersion).
			Configure(envoy_clusters.EdsCluster(service)).
			Configure(envoy_clusters.LbSubset(tagKeySlice)).
			Configure(envoy_clusters.DefaultTimeout()).
			Build()
		if err != nil {
			return nil, err
		}
		resources = append(resources, &model.Resource{
			Name:     service,
			Origin:   OriginIngress,
			Resource: edsCluster,
		})
	}
	return
}

func (_ IngressGenerator) generateEDS(
	proxy *model.Proxy,
	services []string,
	apiVersion envoy_common.APIVersion,
) (resources []*model.Resource, err error) {
	for _, service := range services {
		endpoints := proxy.Routing.OutboundTargets[service]
		cla, err := envoy_endpoints.CreateClusterLoadAssignment(service, endpoints, apiVersion)
		if err != nil {
			return nil, err
		}
		resources = append(resources, &model.Resource{
			Name:     service,
			Origin:   OriginIngress,
			Resource: cla,
		})
	}
	return
}
