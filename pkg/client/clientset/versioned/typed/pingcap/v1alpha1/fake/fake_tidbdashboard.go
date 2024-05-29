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

// FakeTidbDashboards implements TidbDashboardInterface
type FakeTidbDashboards struct {
	Fake *FakePingcapV1alpha1
	ns   string
}

var tidbdashboardsResource = v1alpha1.SchemeGroupVersion.WithResource("tidbdashboards")

var tidbdashboardsKind = v1alpha1.SchemeGroupVersion.WithKind("TidbDashboard")

// Get takes name of the tidbDashboard, and returns the corresponding tidbDashboard object, and an error if there is any.
func (c *FakeTidbDashboards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TidbDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tidbdashboardsResource, c.ns, name), &v1alpha1.TidbDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbDashboard), err
}

// List takes label and field selectors, and returns the list of TidbDashboards that match those selectors.
func (c *FakeTidbDashboards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TidbDashboardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tidbdashboardsResource, tidbdashboardsKind, c.ns, opts), &v1alpha1.TidbDashboardList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TidbDashboardList{ListMeta: obj.(*v1alpha1.TidbDashboardList).ListMeta}
	for _, item := range obj.(*v1alpha1.TidbDashboardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tidbDashboards.
func (c *FakeTidbDashboards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tidbdashboardsResource, c.ns, opts))

}

// Create takes the representation of a tidbDashboard and creates it.  Returns the server's representation of the tidbDashboard, and an error, if there is any.
func (c *FakeTidbDashboards) Create(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.CreateOptions) (result *v1alpha1.TidbDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tidbdashboardsResource, c.ns, tidbDashboard), &v1alpha1.TidbDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbDashboard), err
}

// Update takes the representation of a tidbDashboard and updates it. Returns the server's representation of the tidbDashboard, and an error, if there is any.
func (c *FakeTidbDashboards) Update(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.UpdateOptions) (result *v1alpha1.TidbDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tidbdashboardsResource, c.ns, tidbDashboard), &v1alpha1.TidbDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbDashboard), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTidbDashboards) UpdateStatus(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.UpdateOptions) (*v1alpha1.TidbDashboard, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tidbdashboardsResource, "status", c.ns, tidbDashboard), &v1alpha1.TidbDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbDashboard), err
}

// Delete takes name of the tidbDashboard and deletes it. Returns an error if one occurs.
func (c *FakeTidbDashboards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(tidbdashboardsResource, c.ns, name, opts), &v1alpha1.TidbDashboard{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTidbDashboards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tidbdashboardsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TidbDashboardList{})
	return err
}

// Patch applies the patch and returns the patched tidbDashboard.
func (c *FakeTidbDashboards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TidbDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tidbdashboardsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TidbDashboard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TidbDashboard), err
}
