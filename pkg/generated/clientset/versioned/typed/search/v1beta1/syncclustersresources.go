/*
Copyright The Karbour Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/KusionStack/karbour/pkg/apis/search/v1beta1"
	scheme "github.com/KusionStack/karbour/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SyncClustersResourcesesGetter has a method to return a SyncClustersResourcesInterface.
// A group's client should implement this interface.
type SyncClustersResourcesesGetter interface {
	SyncClustersResourceses() SyncClustersResourcesInterface
}

// SyncClustersResourcesInterface has methods to work with SyncClustersResources resources.
type SyncClustersResourcesInterface interface {
	Create(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.CreateOptions) (*v1beta1.SyncClustersResources, error)
	Update(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.UpdateOptions) (*v1beta1.SyncClustersResources, error)
	UpdateStatus(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.UpdateOptions) (*v1beta1.SyncClustersResources, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.SyncClustersResources, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SyncClustersResourcesList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SyncClustersResources, err error)
	SyncClustersResourcesExpansion
}

// syncClustersResourceses implements SyncClustersResourcesInterface
type syncClustersResourceses struct {
	client rest.Interface
}

// newSyncClustersResourceses returns a SyncClustersResourceses
func newSyncClustersResourceses(c *SearchV1beta1Client) *syncClustersResourceses {
	return &syncClustersResourceses{
		client: c.RESTClient(),
	}
}

// Get takes name of the syncClustersResources, and returns the corresponding syncClustersResources object, and an error if there is any.
func (c *syncClustersResourceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SyncClustersResources, err error) {
	result = &v1beta1.SyncClustersResources{}
	err = c.client.Get().
		Resource("syncclustersresourceses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SyncClustersResourceses that match those selectors.
func (c *syncClustersResourceses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SyncClustersResourcesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SyncClustersResourcesList{}
	err = c.client.Get().
		Resource("syncclustersresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested syncClustersResourceses.
func (c *syncClustersResourceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("syncclustersresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a syncClustersResources and creates it.  Returns the server's representation of the syncClustersResources, and an error, if there is any.
func (c *syncClustersResourceses) Create(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.CreateOptions) (result *v1beta1.SyncClustersResources, err error) {
	result = &v1beta1.SyncClustersResources{}
	err = c.client.Post().
		Resource("syncclustersresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(syncClustersResources).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a syncClustersResources and updates it. Returns the server's representation of the syncClustersResources, and an error, if there is any.
func (c *syncClustersResourceses) Update(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.UpdateOptions) (result *v1beta1.SyncClustersResources, err error) {
	result = &v1beta1.SyncClustersResources{}
	err = c.client.Put().
		Resource("syncclustersresourceses").
		Name(syncClustersResources.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(syncClustersResources).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *syncClustersResourceses) UpdateStatus(ctx context.Context, syncClustersResources *v1beta1.SyncClustersResources, opts v1.UpdateOptions) (result *v1beta1.SyncClustersResources, err error) {
	result = &v1beta1.SyncClustersResources{}
	err = c.client.Put().
		Resource("syncclustersresourceses").
		Name(syncClustersResources.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(syncClustersResources).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the syncClustersResources and deletes it. Returns an error if one occurs.
func (c *syncClustersResourceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("syncclustersresourceses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *syncClustersResourceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("syncclustersresourceses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched syncClustersResources.
func (c *syncClustersResourceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SyncClustersResources, err error) {
	result = &v1beta1.SyncClustersResources{}
	err = c.client.Patch(pt).
		Resource("syncclustersresourceses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}