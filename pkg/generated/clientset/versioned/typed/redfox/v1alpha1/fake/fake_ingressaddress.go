// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/krafton-hq/redfox/pkg/apis/redfox/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressAddresses implements IngressAddressInterface
type FakeIngressAddresses struct {
	Fake *FakeMetadataV1alpha1
	ns   string
}

var ingressaddressesResource = schema.GroupVersionResource{Group: "metadata.sbx-central.io", Version: "v1alpha1", Resource: "ingressaddresses"}

var ingressaddressesKind = schema.GroupVersionKind{Group: "metadata.sbx-central.io", Version: "v1alpha1", Kind: "IngressAddress"}

// Get takes name of the ingressAddress, and returns the corresponding ingressAddress object, and an error if there is any.
func (c *FakeIngressAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IngressAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingressaddressesResource, c.ns, name), &v1alpha1.IngressAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressAddress), err
}

// List takes label and field selectors, and returns the list of IngressAddresses that match those selectors.
func (c *FakeIngressAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IngressAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingressaddressesResource, ingressaddressesKind, c.ns, opts), &v1alpha1.IngressAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IngressAddressList{ListMeta: obj.(*v1alpha1.IngressAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.IngressAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressAddresses.
func (c *FakeIngressAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingressaddressesResource, c.ns, opts))

}

// Create takes the representation of a ingressAddress and creates it.  Returns the server's representation of the ingressAddress, and an error, if there is any.
func (c *FakeIngressAddresses) Create(ctx context.Context, ingressAddress *v1alpha1.IngressAddress, opts v1.CreateOptions) (result *v1alpha1.IngressAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingressaddressesResource, c.ns, ingressAddress), &v1alpha1.IngressAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressAddress), err
}

// Update takes the representation of a ingressAddress and updates it. Returns the server's representation of the ingressAddress, and an error, if there is any.
func (c *FakeIngressAddresses) Update(ctx context.Context, ingressAddress *v1alpha1.IngressAddress, opts v1.UpdateOptions) (result *v1alpha1.IngressAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingressaddressesResource, c.ns, ingressAddress), &v1alpha1.IngressAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressAddress), err
}

// Delete takes name of the ingressAddress and deletes it. Returns an error if one occurs.
func (c *FakeIngressAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ingressaddressesResource, c.ns, name, opts), &v1alpha1.IngressAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingressaddressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IngressAddressList{})
	return err
}

// Patch applies the patch and returns the patched ingressAddress.
func (c *FakeIngressAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IngressAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressaddressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IngressAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IngressAddress), err
}