package inspect_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	gomega_types "github.com/onsi/gomega/types"

	"github.com/spf13/cobra"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/app/kumactl/cmd"
	kumactl_cmd "github.com/kumahq/kuma/app/kumactl/pkg/cmd"
	"github.com/kumahq/kuma/app/kumactl/pkg/resources"
	config_proto "github.com/kumahq/kuma/pkg/config/app/kumactl/v1alpha1"
	mesh_core "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	test_model "github.com/kumahq/kuma/pkg/test/resources/model"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	"github.com/kumahq/kuma/pkg/util/test"
)

type testDataplaneOverviewClient struct {
	receivedTags    map[string]string
	receivedGateway bool
	receivedIngress bool
	total           uint32
	overviews       []*mesh_core.DataplaneOverviewResource
}

func (c *testDataplaneOverviewClient) List(_ context.Context, _ string, tags map[string]string, gateway bool, ingress bool) (*mesh_core.DataplaneOverviewResourceList, error) {
	c.receivedTags = tags
	c.receivedGateway = gateway
	c.receivedIngress = ingress
	return &mesh_core.DataplaneOverviewResourceList{
		Items: c.overviews,
		Pagination: model.Pagination{
			Total: c.total,
		},
	}, nil
}

var _ resources.DataplaneOverviewClient = &testDataplaneOverviewClient{}

var _ = Describe("kumactl inspect dataplanes", func() {

	var now, t1, t2 time.Time
	var sampleDataplaneOverview []*mesh_core.DataplaneOverviewResource

	BeforeEach(func() {
		now, _ = time.Parse(time.RFC3339, "2019-07-17T18:08:41+00:00")
		t1, _ = time.Parse(time.RFC3339, "2018-07-17T16:05:36.995+00:00")
		t2, _ = time.Parse(time.RFC3339, "2019-07-17T16:05:36.995+00:00")
		time.Local = time.UTC

		sampleDataplaneOverview = []*mesh_core.DataplaneOverviewResource{
			{
				Meta: &test_model.ResourceMeta{
					Mesh:             "default",
					Name:             "experiment",
					CreationTime:     t1,
					ModificationTime: now,
				},
				Spec: &mesh_proto.DataplaneOverview{
					Dataplane: &mesh_proto.Dataplane{
						Networking: &mesh_proto.Dataplane_Networking{
							Address: "127.0.0.1",
							Inbound: []*mesh_proto.Dataplane_Networking_Inbound{
								{
									Port:        8080,
									ServicePort: 80,
									Tags: map[string]string{
										mesh_proto.ServiceTag: "mobile",
										"version":             "v1",
									},
								},
								{
									Port:        8090,
									ServicePort: 90,
									Tags: map[string]string{
										mesh_proto.ServiceTag: "metrics",
										"version":             "v1",
									},
								},
							},
						},
					},
					DataplaneInsight: &mesh_proto.DataplaneInsight{
						Subscriptions: []*mesh_proto.DiscoverySubscription{
							{
								Id:                     "1",
								ControlPlaneInstanceId: "node-001",
								ConnectTime:            util_proto.MustTimestampProto(t1),
								Status: &mesh_proto.DiscoverySubscriptionStatus{
									Total: &mesh_proto.DiscoveryServiceStats{
										ResponsesSent:     10,
										ResponsesRejected: 1,
									},
								},
								Version: &mesh_proto.Version{
									KumaDp: &mesh_proto.KumaDpVersion{
										Version:   "1.0.0",
										GitTag:    "v1.0.0",
										GitCommit: "91ce236824a9d875601679aa80c63783fb0e8725",
										BuildDate: "2019-08-07T11:26:06Z",
									},
									Envoy: &mesh_proto.EnvoyVersion{
										Version: "1.16.0",
										Build:   "hash/1.16.0/RELEASE",
									},
								},
							},
							{
								Id:                     "2",
								ControlPlaneInstanceId: "node-002",
								ConnectTime:            util_proto.MustTimestampProto(t2),
								Status: &mesh_proto.DiscoverySubscriptionStatus{
									Total: &mesh_proto.DiscoveryServiceStats{
										ResponsesSent:     20,
										ResponsesRejected: 2,
									},
								},
								Version: &mesh_proto.Version{
									KumaDp: &mesh_proto.KumaDpVersion{
										Version:   "1.0.2",
										GitTag:    "v1.0.2",
										GitCommit: "9d868cd8681c4021bb4a10bf2306ca613ba4de42",
										BuildDate: "2020-08-07T11:26:06Z",
									},
									Envoy: &mesh_proto.EnvoyVersion{
										Version: "1.16.1",
										Build:   "hash/1.16.1/RELEASE",
									},
								},
							},
						},
						MTLS: &mesh_proto.DataplaneInsight_MTLS{
							CertificateExpirationTime: &timestamppb.Timestamp{
								Seconds: 1588926502,
							},
							LastCertificateRegeneration: &timestamppb.Timestamp{
								Seconds: 1563306488,
							},
							CertificateRegenerations: 10,
						},
					},
				},
			},
			{
				Meta: &test_model.ResourceMeta{
					Mesh:             "default",
					Name:             "degraded-dp",
					CreationTime:     t1,
					ModificationTime: now,
				},
				Spec: &mesh_proto.DataplaneOverview{
					Dataplane: &mesh_proto.Dataplane{
						Networking: &mesh_proto.Dataplane_Networking{
							Address: "127.0.0.1",
							Inbound: []*mesh_proto.Dataplane_Networking_Inbound{
								{
									Port:        8080,
									ServicePort: 80,
									Tags: map[string]string{
										mesh_proto.ServiceTag: "example",
									},
								},
								{
									Port:        9001,
									ServicePort: 81,
									Health:      &mesh_proto.Dataplane_Networking_Inbound_Health{Ready: false},
									Tags: map[string]string{
										mesh_proto.ServiceTag: "example",
									},
								},
							},
						},
					},
					DataplaneInsight: &mesh_proto.DataplaneInsight{
						Subscriptions: []*mesh_proto.DiscoverySubscription{
							{
								Id:                     "1",
								ControlPlaneInstanceId: "node-001",
								ConnectTime:            util_proto.MustTimestampProto(t1),
								Status: &mesh_proto.DiscoverySubscriptionStatus{
									Total: &mesh_proto.DiscoveryServiceStats{
										ResponsesSent:     10,
										ResponsesRejected: 1,
									},
								},
								Version: &mesh_proto.Version{
									KumaDp: &mesh_proto.KumaDpVersion{
										Version:   "1.0.0",
										GitTag:    "v1.0.0",
										GitCommit: "91ce236824a9d875601679aa80c63783fb0e8725",
										BuildDate: "2019-08-07T11:26:06Z",
									},
									Envoy: &mesh_proto.EnvoyVersion{
										Version: "1.16.0",
										Build:   "hash/1.16.0/RELEASE",
									},
								},
							},
							{
								Id:                     "2",
								ControlPlaneInstanceId: "node-002",
								ConnectTime:            util_proto.MustTimestampProto(t2),
								Status: &mesh_proto.DiscoverySubscriptionStatus{
									Total: &mesh_proto.DiscoveryServiceStats{
										ResponsesSent:     20,
										ResponsesRejected: 2,
									},
								},
								Version: &mesh_proto.Version{
									KumaDp: &mesh_proto.KumaDpVersion{
										Version:   "1.0.2",
										GitTag:    "v1.0.2",
										GitCommit: "9d868cd8681c4021bb4a10bf2306ca613ba4de42",
										BuildDate: "2020-08-07T11:26:06Z",
									},
									Envoy: &mesh_proto.EnvoyVersion{
										Version: "1.16.1",
										Build:   "hash/1.16.1/RELEASE",
									},
								},
							},
						},
						MTLS: &mesh_proto.DataplaneInsight_MTLS{
							CertificateExpirationTime: &timestamppb.Timestamp{
								Seconds: 1588926502,
							},
							LastCertificateRegeneration: &timestamppb.Timestamp{
								Seconds: 1563306488,
							},
							CertificateRegenerations: 10,
						},
					},
				},
			},
			{
				Meta: &test_model.ResourceMeta{
					Mesh:             "default",
					Name:             "offline-dp",
					CreationTime:     t1,
					ModificationTime: now,
				},
				Spec: &mesh_proto.DataplaneOverview{
					Dataplane: &mesh_proto.Dataplane{
						Networking: &mesh_proto.Dataplane_Networking{
							Address: "127.0.0.1",
							Inbound: []*mesh_proto.Dataplane_Networking_Inbound{
								{
									Port:        8080,
									ServicePort: 80,
									Health:      &mesh_proto.Dataplane_Networking_Inbound_Health{Ready: false},
									Tags: map[string]string{
										mesh_proto.ServiceTag: "example",
									},
								},
								{
									Port:        9001,
									ServicePort: 81,
									Health:      &mesh_proto.Dataplane_Networking_Inbound_Health{Ready: false},
									Tags: map[string]string{
										mesh_proto.ServiceTag: "example",
									},
								},
							},
						},
					},
					DataplaneInsight: &mesh_proto.DataplaneInsight{
						Subscriptions: []*mesh_proto.DiscoverySubscription{
							{
								Id:                     "1",
								ControlPlaneInstanceId: "node-001",
								ConnectTime:            util_proto.MustTimestampProto(t1),
								Status: &mesh_proto.DiscoverySubscriptionStatus{
									Total: &mesh_proto.DiscoveryServiceStats{
										ResponsesSent:     10,
										ResponsesRejected: 1,
									},
								},
								Version: &mesh_proto.Version{
									KumaDp: &mesh_proto.KumaDpVersion{
										Version:   "1.0.0",
										GitTag:    "v1.0.0",
										GitCommit: "91ce236824a9d875601679aa80c63783fb0e8725",
										BuildDate: "2019-08-07T11:26:06Z",
									},
									Envoy: &mesh_proto.EnvoyVersion{
										Version: "1.16.0",
										Build:   "hash/1.16.0/RELEASE",
									},
								},
							},
							{
								Id:                     "2",
								ControlPlaneInstanceId: "node-002",
								ConnectTime:            util_proto.MustTimestampProto(t2),
								Status: &mesh_proto.DiscoverySubscriptionStatus{
									Total: &mesh_proto.DiscoveryServiceStats{
										ResponsesSent:     20,
										ResponsesRejected: 2,
									},
								},
								Version: &mesh_proto.Version{
									KumaDp: &mesh_proto.KumaDpVersion{
										Version:   "1.0.2",
										GitTag:    "v1.0.2",
										GitCommit: "9d868cd8681c4021bb4a10bf2306ca613ba4de42",
										BuildDate: "2020-08-07T11:26:06Z",
									},
									Envoy: &mesh_proto.EnvoyVersion{
										Version: "1.16.1",
										Build:   "hash/1.16.1/RELEASE",
									},
								},
							},
						},
						MTLS: &mesh_proto.DataplaneInsight_MTLS{
							CertificateExpirationTime: &timestamppb.Timestamp{
								Seconds: 1588926502,
							},
							LastCertificateRegeneration: &timestamppb.Timestamp{
								Seconds: 1563306488,
							},
							CertificateRegenerations: 10,
						},
					},
				},
			},
			{
				Meta: &test_model.ResourceMeta{
					Mesh:             "default",
					Name:             "example",
					CreationTime:     t1,
					ModificationTime: now,
				},
				Spec: &mesh_proto.DataplaneOverview{
					Dataplane: &mesh_proto.Dataplane{
						Networking: &mesh_proto.Dataplane_Networking{
							Address: "127.0.0.1",
							Inbound: []*mesh_proto.Dataplane_Networking_Inbound{
								{
									Port:        8080,
									ServicePort: 80,
									Tags: map[string]string{
										"kuma.io/service": "example",
									},
								},
							},
						},
					},
					DataplaneInsight: &mesh_proto.DataplaneInsight{
						Subscriptions: []*mesh_proto.DiscoverySubscription{
							{
								Id:                     "1",
								ControlPlaneInstanceId: "node-001",
							},
							{
								Id:                     "2",
								ControlPlaneInstanceId: "node-002",
							},
							{
								Id:                     "3",
								ControlPlaneInstanceId: "node-003",
							},
						},
					},
				},
			},
		}
	})

	Describe("InspectDataplanesCmd", func() {

		var rootCtx *kumactl_cmd.RootContext
		var rootCmd *cobra.Command
		var buf *bytes.Buffer

		var testClient *testDataplaneOverviewClient

		BeforeEach(func() {
			// setup
			testClient = &testDataplaneOverviewClient{
				total:     uint32(len(sampleDataplaneOverview)),
				overviews: sampleDataplaneOverview,
			}

			rootCtx = &kumactl_cmd.RootContext{
				Runtime: kumactl_cmd.RootRuntime{
					Now: func() time.Time { return now },
					NewDataplaneOverviewClient: func(*config_proto.ControlPlaneCoordinates_ApiServer) (resources.DataplaneOverviewClient, error) {
						return testClient, nil
					},
					NewAPIServerClient: test.GetMockNewAPIServerClient(),
				},
			}

			rootCmd = cmd.NewRootCmd(rootCtx)
			buf = &bytes.Buffer{}
			rootCmd.SetOut(buf)
		})

		type testCase struct {
			outputFormat string
			goldenFile   string
			matcher      func(interface{}) gomega_types.GomegaMatcher
		}

		byLine := func(s string) []string {
			lines := strings.Split(s, "\n")
			var trimmedLines []string
			for _, line := range lines {
				trimmedLines = append(trimmedLines, strings.TrimSpace(line))
			}
			return trimmedLines
		}

		DescribeTable("kumactl inspect dataplanes -o table|json|yaml",
			func(given testCase) {
				// given
				rootCmd.SetArgs(append([]string{
					"--config-file", filepath.Join("..", "testdata", "sample-kumactl.config.yaml"),
					"inspect", "dataplanes"}, given.outputFormat))

				// when
				err := rootCmd.Execute()
				// then
				Expect(err).ToNot(HaveOccurred())

				// when
				expected, err := ioutil.ReadFile(filepath.Join("testdata", given.goldenFile))
				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(buf.String()).To(given.matcher(expected))
			},
			Entry("should support Table output by default", testCase{
				outputFormat: "",
				goldenFile:   "inspect-dataplanes.golden.txt",
				matcher: func(expected interface{}) gomega_types.GomegaMatcher {
					return WithTransform(byLine, Equal(byLine(string(expected.([]byte)))))
				},
			}),
			Entry("should support Table output explicitly", testCase{
				outputFormat: "-otable",
				goldenFile:   "inspect-dataplanes.golden.txt",
				matcher: func(expected interface{}) gomega_types.GomegaMatcher {
					return WithTransform(byLine, Equal(byLine(string(expected.([]byte)))))
				},
			}),
			Entry("should support JSON output", testCase{
				outputFormat: "-ojson",
				goldenFile:   "inspect-dataplanes.golden.json",
				matcher:      MatchJSON,
			}),
			Entry("should support YAML output", testCase{
				outputFormat: "-oyaml",
				goldenFile:   "inspect-dataplanes.golden.yaml",
				matcher:      MatchYAML,
			}),
		)

		Describe("kumactl inspect dataplanes --tag", func() {
			It("tags should be passed to the client", func() {
				// given
				rootCmd.SetArgs([]string{
					"--config-file", filepath.Join("..", "testdata", "sample-kumactl.config.yaml"),
					"inspect", "dataplanes", "--tag", "kuma.io/service=mobile", "--tag", "version=v1"})

				// when
				err := rootCmd.Execute()
				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(testClient.receivedTags).To(HaveKeyWithValue(mesh_proto.ServiceTag, "mobile"))
				Expect(testClient.receivedTags).To(HaveKeyWithValue("version", "v1"))
			})
		})

		Describe("kumactl inspect dataplanes --gateway", func() {
			It("gateway should be passed to the client", func() {
				// given
				rootCmd.SetArgs([]string{
					"--config-file", filepath.Join("..", "testdata", "sample-kumactl.config.yaml"),
					"inspect", "dataplanes", "--gateway"})

				// when
				err := rootCmd.Execute()
				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(testClient.receivedGateway).To(BeTrue())
			})
		})

		Describe("kumactl inspect dataplanes --ingress", func() {
			It("gateway should be passed to the client", func() {
				// given
				rootCmd.SetArgs([]string{
					"--config-file", filepath.Join("..", "testdata", "sample-kumactl.config.yaml"),
					"inspect", "dataplanes", "--ingress"})

				// when
				err := rootCmd.Execute()
				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(testClient.receivedIngress).To(BeTrue())
			})
		})
	})
})
