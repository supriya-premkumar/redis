/*
Copyright The KubeDB Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisOpsRequestsGetter has a method to return a RedisOpsRequestInterface.
// A group's client should implement this interface.
type RedisOpsRequestsGetter interface {
	RedisOpsRequests(namespace string) RedisOpsRequestInterface
}

// RedisOpsRequestInterface has methods to work with RedisOpsRequest resources.
type RedisOpsRequestInterface interface {
	Create(*v1alpha1.RedisOpsRequest) (*v1alpha1.RedisOpsRequest, error)
	Update(*v1alpha1.RedisOpsRequest) (*v1alpha1.RedisOpsRequest, error)
	UpdateStatus(*v1alpha1.RedisOpsRequest) (*v1alpha1.RedisOpsRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RedisOpsRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.RedisOpsRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisOpsRequest, err error)
	RedisOpsRequestExpansion
}

// redisOpsRequests implements RedisOpsRequestInterface
type redisOpsRequests struct {
	client rest.Interface
	ns     string
}

// newRedisOpsRequests returns a RedisOpsRequests
func newRedisOpsRequests(c *OpsV1alpha1Client, namespace string) *redisOpsRequests {
	return &redisOpsRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisOpsRequest, and returns the corresponding redisOpsRequest object, and an error if there is any.
func (c *redisOpsRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisOpsRequest, err error) {
	result = &v1alpha1.RedisOpsRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisopsrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisOpsRequests that match those selectors.
func (c *redisOpsRequests) List(opts v1.ListOptions) (result *v1alpha1.RedisOpsRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedisOpsRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisopsrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisOpsRequests.
func (c *redisOpsRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisopsrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a redisOpsRequest and creates it.  Returns the server's representation of the redisOpsRequest, and an error, if there is any.
func (c *redisOpsRequests) Create(redisOpsRequest *v1alpha1.RedisOpsRequest) (result *v1alpha1.RedisOpsRequest, err error) {
	result = &v1alpha1.RedisOpsRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisopsrequests").
		Body(redisOpsRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisOpsRequest and updates it. Returns the server's representation of the redisOpsRequest, and an error, if there is any.
func (c *redisOpsRequests) Update(redisOpsRequest *v1alpha1.RedisOpsRequest) (result *v1alpha1.RedisOpsRequest, err error) {
	result = &v1alpha1.RedisOpsRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisopsrequests").
		Name(redisOpsRequest.Name).
		Body(redisOpsRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *redisOpsRequests) UpdateStatus(redisOpsRequest *v1alpha1.RedisOpsRequest) (result *v1alpha1.RedisOpsRequest, err error) {
	result = &v1alpha1.RedisOpsRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisopsrequests").
		Name(redisOpsRequest.Name).
		SubResource("status").
		Body(redisOpsRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisOpsRequest and deletes it. Returns an error if one occurs.
func (c *redisOpsRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisopsrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisOpsRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisopsrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisOpsRequest.
func (c *redisOpsRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisOpsRequest, err error) {
	result = &v1alpha1.RedisOpsRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisopsrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
