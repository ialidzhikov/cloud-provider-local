package local

import (
	cloudprovider "k8s.io/cloud-provider"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type LocalCloud struct {
	SourceClient client.Client
}

var _ cloudprovider.Interface = &LocalCloud{}

func (c *LocalCloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {

}

func (c *LocalCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

func (c *LocalCloud) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

func (c *LocalCloud) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return &LocalInstancesV2{client: c.SourceClient}, true
}

func (c *LocalCloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

func (c *LocalCloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

func (c *LocalCloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

func (c *LocalCloud) ProviderName() string {
	return ""
}

func (c *LocalCloud) HasClusterID() bool {
	return false
}
