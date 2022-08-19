// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/krafton-hq/red-fox/pkg/apis/redfox/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNatIps implements NatIpInterface
type FakeNatIps struct {
	Fake *FakeMetadataV1alpha1
	ns   string
}

var natipsResource = schema.GroupVersionResource{Group: "metadata.sbx-central.io", Version: "v1alpha1", Resource: "natips"}

var natipsKind = schema.GroupVersionKind{Group: "metadata.sbx-central.io", Version: "v1alpha1", Kind: "NatIp"}

// Get takes name of the natIp, and returns the corresponding natIp object, and an error if there is any.
func (c *FakeNatIps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NatIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(natipsResource, c.ns, name), &v1alpha1.NatIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatIp), err
}

// List takes label and field selectors, and returns the list of NatIps that match those selectors.
func (c *FakeNatIps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NatIpList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(natipsResource, natipsKind, c.ns, opts), &v1alpha1.NatIpList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NatIpList{ListMeta: obj.(*v1alpha1.NatIpList).ListMeta}
	for _, item := range obj.(*v1alpha1.NatIpList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested natIps.
func (c *FakeNatIps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(natipsResource, c.ns, opts))

}

// Create takes the representation of a natIp and creates it.  Returns the server's representation of the natIp, and an error, if there is any.
func (c *FakeNatIps) Create(ctx context.Context, natIp *v1alpha1.NatIp, opts v1.CreateOptions) (result *v1alpha1.NatIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(natipsResource, c.ns, natIp), &v1alpha1.NatIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatIp), err
}

// Update takes the representation of a natIp and updates it. Returns the server's representation of the natIp, and an error, if there is any.
func (c *FakeNatIps) Update(ctx context.Context, natIp *v1alpha1.NatIp, opts v1.UpdateOptions) (result *v1alpha1.NatIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(natipsResource, c.ns, natIp), &v1alpha1.NatIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatIp), err
}

// Delete takes name of the natIp and deletes it. Returns an error if one occurs.
func (c *FakeNatIps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(natipsResource, c.ns, name, opts), &v1alpha1.NatIp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNatIps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(natipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NatIpList{})
	return err
}

// Patch applies the patch and returns the patched natIp.
func (c *FakeNatIps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NatIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(natipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NatIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatIp), err
}
