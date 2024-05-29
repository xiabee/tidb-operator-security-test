// Copyright PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDMClusters implements DMClusterInterface
type FakeDMClusters struct {
	Fake *FakePingcapV1alpha1
	ns   string
}

var dmclustersResource = v1alpha1.SchemeGroupVersion.WithResource("dmclusters")

var dmclustersKind = v1alpha1.SchemeGroupVersion.WithKind("DMCluster")

// Get takes name of the dMCluster, and returns the corresponding dMCluster object, and an error if there is any.
func (c *FakeDMClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DMCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dmclustersResource, c.ns, name), &v1alpha1.DMCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMCluster), err
}

// List takes label and field selectors, and returns the list of DMClusters that match those selectors.
func (c *FakeDMClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DMClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dmclustersResource, dmclustersKind, c.ns, opts), &v1alpha1.DMClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DMClusterList{ListMeta: obj.(*v1alpha1.DMClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DMClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dMClusters.
func (c *FakeDMClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dmclustersResource, c.ns, opts))

}

// Create takes the representation of a dMCluster and creates it.  Returns the server's representation of the dMCluster, and an error, if there is any.
func (c *FakeDMClusters) Create(ctx context.Context, dMCluster *v1alpha1.DMCluster, opts v1.CreateOptions) (result *v1alpha1.DMCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dmclustersResource, c.ns, dMCluster), &v1alpha1.DMCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMCluster), err
}

// Update takes the representation of a dMCluster and updates it. Returns the server's representation of the dMCluster, and an error, if there is any.
func (c *FakeDMClusters) Update(ctx context.Context, dMCluster *v1alpha1.DMCluster, opts v1.UpdateOptions) (result *v1alpha1.DMCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dmclustersResource, c.ns, dMCluster), &v1alpha1.DMCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDMClusters) UpdateStatus(ctx context.Context, dMCluster *v1alpha1.DMCluster, opts v1.UpdateOptions) (*v1alpha1.DMCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dmclustersResource, "status", c.ns, dMCluster), &v1alpha1.DMCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMCluster), err
}

// Delete takes name of the dMCluster and deletes it. Returns an error if one occurs.
func (c *FakeDMClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dmclustersResource, c.ns, name, opts), &v1alpha1.DMCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDMClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dmclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DMClusterList{})
	return err
}

// Patch applies the patch and returns the patched dMCluster.
func (c *FakeDMClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DMCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dmclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DMCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DMCluster), err
}
