package local

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cloudprovider "k8s.io/cloud-provider"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ cloudprovider.InstancesV2 = &LocalInstancesV2{}

type LocalInstancesV2 struct {
	client client.Client
}

func (l LocalInstancesV2) InstanceExists(ctx context.Context, node *v1.Node) (bool, error) {
	nodePod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: node.Name,
			// TODO: Replace this with the namespace that the CCM runs in
			Namespace: "shoot--local--local",
		},
	}
	if err := l.client.Get(ctx, client.ObjectKeyFromObject(nodePod), nodePod); err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (l LocalInstancesV2) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	// A terminated Node Pod is already removed from store.
	// Hence, for now we always return that the instace is not shut-down.
	// We should not return true when the Node Pod does not exists.
	// See https://github.com/kubernetes/cloud-provider-aws/blob/e58b027f08820da1f3b523b5238adbcd4e7e85a2/pkg/providers/v1/aws.go#L943-L944
	return false, nil
}

func (l LocalInstancesV2) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	return &cloudprovider.InstanceMetadata{
		ProviderID:       node.Name,
		InstanceType:     "",
		NodeAddresses:    nil,
		Zone:             "0",
		Region:           "local",
		AdditionalLabels: nil,
	}, nil
}
