// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/container/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainerNodePools implements ContainerNodePoolInterface
type FakeContainerNodePools struct {
	Fake *FakeContainerV1beta1
	ns   string
}

var containernodepoolsResource = schema.GroupVersionResource{Group: "container.cnrm.cloud.google.com", Version: "v1beta1", Resource: "containernodepools"}

var containernodepoolsKind = schema.GroupVersionKind{Group: "container.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ContainerNodePool"}

// Get takes name of the containerNodePool, and returns the corresponding containerNodePool object, and an error if there is any.
func (c *FakeContainerNodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containernodepoolsResource, c.ns, name), &v1beta1.ContainerNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerNodePool), err
}

// List takes label and field selectors, and returns the list of ContainerNodePools that match those selectors.
func (c *FakeContainerNodePools) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ContainerNodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(containernodepoolsResource, containernodepoolsKind, c.ns, opts), &v1beta1.ContainerNodePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ContainerNodePoolList{ListMeta: obj.(*v1beta1.ContainerNodePoolList).ListMeta}
	for _, item := range obj.(*v1beta1.ContainerNodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerNodePools.
func (c *FakeContainerNodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containernodepoolsResource, c.ns, opts))

}

// Create takes the representation of a containerNodePool and creates it.  Returns the server's representation of the containerNodePool, and an error, if there is any.
func (c *FakeContainerNodePools) Create(ctx context.Context, containerNodePool *v1beta1.ContainerNodePool, opts v1.CreateOptions) (result *v1beta1.ContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(containernodepoolsResource, c.ns, containerNodePool), &v1beta1.ContainerNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerNodePool), err
}

// Update takes the representation of a containerNodePool and updates it. Returns the server's representation of the containerNodePool, and an error, if there is any.
func (c *FakeContainerNodePools) Update(ctx context.Context, containerNodePool *v1beta1.ContainerNodePool, opts v1.UpdateOptions) (result *v1beta1.ContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(containernodepoolsResource, c.ns, containerNodePool), &v1beta1.ContainerNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerNodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerNodePools) UpdateStatus(ctx context.Context, containerNodePool *v1beta1.ContainerNodePool, opts v1.UpdateOptions) (*v1beta1.ContainerNodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(containernodepoolsResource, "status", c.ns, containerNodePool), &v1beta1.ContainerNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerNodePool), err
}

// Delete takes name of the containerNodePool and deletes it. Returns an error if one occurs.
func (c *FakeContainerNodePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(containernodepoolsResource, c.ns, name, opts), &v1beta1.ContainerNodePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerNodePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(containernodepoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ContainerNodePoolList{})
	return err
}

// Patch applies the patch and returns the patched containerNodePool.
func (c *FakeContainerNodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ContainerNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(containernodepoolsResource, c.ns, name, pt, data, subresources...), &v1beta1.ContainerNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ContainerNodePool), err
}