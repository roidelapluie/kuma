package ratelimits

import (
	"context"

	"google.golang.org/protobuf/proto"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"

	"github.com/pkg/errors"

	manager_dataplane "github.com/kumahq/kuma/pkg/core/managers/apis/dataplane"
	"github.com/kumahq/kuma/pkg/core/policy"
	mesh_core "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	"github.com/kumahq/kuma/pkg/core/resources/store"
	core_xds "github.com/kumahq/kuma/pkg/core/xds"
)

type RateLimitMatcher struct {
	ResourceManager manager.ReadOnlyResourceManager
}

func (m *RateLimitMatcher) Match(ctx context.Context, dataplane *mesh_core.DataplaneResource, mesh *mesh_core.MeshResource) (core_xds.RateLimitsMap, error) {
	ratelimits := &mesh_core.RateLimitResourceList{}
	if err := m.ResourceManager.List(ctx, ratelimits, store.ListByMesh(dataplane.GetMeta().GetMesh())); err != nil {
		return nil, errors.Wrap(err, "could not retrieve ratelimits")
	}

	return buildRateLimitMap(dataplane, mesh, splitPoliciesBySourceMatch(ratelimits.Items))
}

func buildRateLimitMap(
	dataplane *mesh_core.DataplaneResource,
	mesh *mesh_core.MeshResource,
	rateLimits []*mesh_core.RateLimitResource,
) (core_xds.RateLimitsMap, error) {
	policies := make([]policy.ConnectionPolicy, len(rateLimits))
	for i, ratelimit := range rateLimits {
		policies[i] = ratelimit
	}

	additionalInbounds, err := manager_dataplane.AdditionalInbounds(dataplane, mesh)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch additional inbounds")
	}
	inbounds := append(dataplane.Spec.GetNetworking().GetInbound(), additionalInbounds...)
	policyMap := policy.SelectInboundConnectionMatchingPolicies(dataplane, inbounds, policies)

	result := core_xds.RateLimitsMap{}
	for inbound, connectionPolicies := range policyMap {
		result[inbound] = []*mesh_proto.RateLimit{}
		for _, policy := range connectionPolicies {
			result[inbound] = append(result[inbound], policy.(*mesh_core.RateLimitResource).Spec)
		}
	}
	return result, nil
}

// We need to split policies with many sources into individual policies, because otherwise
// we would not be able to sort all selectors from all policies in the most specific to the
// most general order.
//
// Example:
//  sources:
//    - match:
//        kuma.io/service: 'web'
//        version: '1'
//    - match:
//        kuma.io/service: 'web-api'
//---
//  sources:
//    - match:
//        kuma.io/service: 'web'
//    - match:
//        kuma.io/service: 'web-api'
//        version: '1'
func splitPoliciesBySourceMatch(rateLimits []*mesh_core.RateLimitResource) []*mesh_core.RateLimitResource {
	result := []*mesh_core.RateLimitResource{}

	for _, rateLimit := range rateLimits {
		for i := range rateLimit.Sources() {
			newRateLimit := &mesh_core.RateLimitResource{
				Meta: rateLimit.GetMeta(),
				Spec: proto.Clone(rateLimit.Spec).(*mesh_proto.RateLimit),
			}
			newRateLimit.Spec.Sources = []*mesh_proto.Selector{newRateLimit.Spec.Sources[i]}

			result = append(result, newRateLimit)
		}
	}

	return result
}
