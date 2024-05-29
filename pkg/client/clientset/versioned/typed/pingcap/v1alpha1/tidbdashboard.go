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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	scheme "github.com/pingcap/tidb-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TidbDashboardsGetter has a method to return a TidbDashboardInterface.
// A group's client should implement this interface.
type TidbDashboardsGetter interface {
	TidbDashboards(namespace string) TidbDashboardInterface
}

// TidbDashboardInterface has methods to work with TidbDashboard resources.
type TidbDashboardInterface interface {
	Create(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.CreateOptions) (*v1alpha1.TidbDashboard, error)
	Update(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.UpdateOptions) (*v1alpha1.TidbDashboard, error)
	UpdateStatus(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.UpdateOptions) (*v1alpha1.TidbDashboard, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TidbDashboard, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TidbDashboardList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TidbDashboard, err error)
	TidbDashboardExpansion
}

// tidbDashboards implements TidbDashboardInterface
type tidbDashboards struct {
	client rest.Interface
	ns     string
}

// newTidbDashboards returns a TidbDashboards
func newTidbDashboards(c *PingcapV1alpha1Client, namespace string) *tidbDashboards {
	return &tidbDashboards{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tidbDashboard, and returns the corresponding tidbDashboard object, and an error if there is any.
func (c *tidbDashboards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TidbDashboard, err error) {
	result = &v1alpha1.TidbDashboard{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tidbdashboards").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TidbDashboards that match those selectors.
func (c *tidbDashboards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TidbDashboardList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TidbDashboardList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tidbdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tidbDashboards.
func (c *tidbDashboards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tidbdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tidbDashboard and creates it.  Returns the server's representation of the tidbDashboard, and an error, if there is any.
func (c *tidbDashboards) Create(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.CreateOptions) (result *v1alpha1.TidbDashboard, err error) {
	result = &v1alpha1.TidbDashboard{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tidbdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tidbDashboard).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tidbDashboard and updates it. Returns the server's representation of the tidbDashboard, and an error, if there is any.
func (c *tidbDashboards) Update(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.UpdateOptions) (result *v1alpha1.TidbDashboard, err error) {
	result = &v1alpha1.TidbDashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tidbdashboards").
		Name(tidbDashboard.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tidbDashboard).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tidbDashboards) UpdateStatus(ctx context.Context, tidbDashboard *v1alpha1.TidbDashboard, opts v1.UpdateOptions) (result *v1alpha1.TidbDashboard, err error) {
	result = &v1alpha1.TidbDashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tidbdashboards").
		Name(tidbDashboard.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tidbDashboard).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tidbDashboard and deletes it. Returns an error if one occurs.
func (c *tidbDashboards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tidbdashboards").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tidbDashboards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tidbdashboards").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tidbDashboard.
func (c *tidbDashboards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TidbDashboard, err error) {
	result = &v1alpha1.TidbDashboard{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tidbdashboards").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}