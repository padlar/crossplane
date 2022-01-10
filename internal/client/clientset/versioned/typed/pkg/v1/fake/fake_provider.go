/*
Copyright 2021 The Crossplane Authors.

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

package fake

import (
	"context"

	pkgv1 "github.com/crossplane/crossplane/apis/pkg/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProviders implements ProviderInterface
type FakeProviders struct {
	Fake *FakePkgV1
}

var providersResource = schema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "providers"}

var providersKind = schema.GroupVersionKind{Group: "pkg.crossplane.io", Version: "v1", Kind: "Provider"}

// Get takes name of the provider, and returns the corresponding provider object, and an error if there is any.
func (c *FakeProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *pkgv1.Provider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(providersResource, name), &pkgv1.Provider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Provider), err
}

// List takes label and field selectors, and returns the list of Providers that match those selectors.
func (c *FakeProviders) List(ctx context.Context, opts v1.ListOptions) (result *pkgv1.ProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(providersResource, providersKind, opts), &pkgv1.ProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pkgv1.ProviderList{ListMeta: obj.(*pkgv1.ProviderList).ListMeta}
	for _, item := range obj.(*pkgv1.ProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested providers.
func (c *FakeProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(providersResource, opts))
}

// Create takes the representation of a provider and creates it.  Returns the server's representation of the provider, and an error, if there is any.
func (c *FakeProviders) Create(ctx context.Context, provider *pkgv1.Provider, opts v1.CreateOptions) (result *pkgv1.Provider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(providersResource, provider), &pkgv1.Provider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Provider), err
}

// Update takes the representation of a provider and updates it. Returns the server's representation of the provider, and an error, if there is any.
func (c *FakeProviders) Update(ctx context.Context, provider *pkgv1.Provider, opts v1.UpdateOptions) (result *pkgv1.Provider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(providersResource, provider), &pkgv1.Provider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Provider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProviders) UpdateStatus(ctx context.Context, provider *pkgv1.Provider, opts v1.UpdateOptions) (*pkgv1.Provider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(providersResource, "status", provider), &pkgv1.Provider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Provider), err
}

// Delete takes name of the provider and deletes it. Returns an error if one occurs.
func (c *FakeProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(providersResource, name, opts), &pkgv1.Provider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(providersResource, listOpts)

	_, err := c.Fake.Invokes(action, &pkgv1.ProviderList{})
	return err
}

// Patch applies the patch and returns the patched provider.
func (c *FakeProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *pkgv1.Provider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(providersResource, name, pt, data, subresources...), &pkgv1.Provider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Provider), err
}
