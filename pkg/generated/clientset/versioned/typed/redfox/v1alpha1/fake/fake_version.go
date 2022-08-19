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

// FakeVersions implements VersionInterface
type FakeVersions struct {
	Fake *FakeMetadataV1alpha1
	ns   string
}

var versionsResource = schema.GroupVersionResource{Group: "metadata.sbx-central.io", Version: "v1alpha1", Resource: "versions"}

var versionsKind = schema.GroupVersionKind{Group: "metadata.sbx-central.io", Version: "v1alpha1", Kind: "Version"}

// Get takes name of the version, and returns the corresponding version object, and an error if there is any.
func (c *FakeVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(versionsResource, c.ns, name), &v1alpha1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Version), err
}

// List takes label and field selectors, and returns the list of Versions that match those selectors.
func (c *FakeVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(versionsResource, versionsKind, c.ns, opts), &v1alpha1.VersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VersionList{ListMeta: obj.(*v1alpha1.VersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.VersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested versions.
func (c *FakeVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(versionsResource, c.ns, opts))

}

// Create takes the representation of a version and creates it.  Returns the server's representation of the version, and an error, if there is any.
func (c *FakeVersions) Create(ctx context.Context, version *v1alpha1.Version, opts v1.CreateOptions) (result *v1alpha1.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(versionsResource, c.ns, version), &v1alpha1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Version), err
}

// Update takes the representation of a version and updates it. Returns the server's representation of the version, and an error, if there is any.
func (c *FakeVersions) Update(ctx context.Context, version *v1alpha1.Version, opts v1.UpdateOptions) (result *v1alpha1.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(versionsResource, c.ns, version), &v1alpha1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Version), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVersions) UpdateStatus(ctx context.Context, version *v1alpha1.Version, opts v1.UpdateOptions) (*v1alpha1.Version, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(versionsResource, "status", c.ns, version), &v1alpha1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Version), err
}

// Delete takes name of the version and deletes it. Returns an error if one occurs.
func (c *FakeVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(versionsResource, c.ns, name, opts), &v1alpha1.Version{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(versionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VersionList{})
	return err
}

// Patch applies the patch and returns the patched version.
func (c *FakeVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Version, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(versionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Version), err
}
