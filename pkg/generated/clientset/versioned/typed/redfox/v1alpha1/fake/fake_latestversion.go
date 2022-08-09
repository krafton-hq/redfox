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

// FakeLatestVersions implements LatestVersionInterface
type FakeLatestVersions struct {
	Fake *FakeMetadataV1alpha1
	ns   string
}

var latestversionsResource = schema.GroupVersionResource{Group: "metadata.sbx-central.io", Version: "v1alpha1", Resource: "latestversions"}

var latestversionsKind = schema.GroupVersionKind{Group: "metadata.sbx-central.io", Version: "v1alpha1", Kind: "LatestVersion"}

// Get takes name of the latestVersion, and returns the corresponding latestVersion object, and an error if there is any.
func (c *FakeLatestVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LatestVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(latestversionsResource, c.ns, name), &v1alpha1.LatestVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LatestVersion), err
}

// List takes label and field selectors, and returns the list of LatestVersions that match those selectors.
func (c *FakeLatestVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LatestVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(latestversionsResource, latestversionsKind, c.ns, opts), &v1alpha1.LatestVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LatestVersionList{ListMeta: obj.(*v1alpha1.LatestVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.LatestVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested latestVersions.
func (c *FakeLatestVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(latestversionsResource, c.ns, opts))

}

// Create takes the representation of a latestVersion and creates it.  Returns the server's representation of the latestVersion, and an error, if there is any.
func (c *FakeLatestVersions) Create(ctx context.Context, latestVersion *v1alpha1.LatestVersion, opts v1.CreateOptions) (result *v1alpha1.LatestVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(latestversionsResource, c.ns, latestVersion), &v1alpha1.LatestVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LatestVersion), err
}

// Update takes the representation of a latestVersion and updates it. Returns the server's representation of the latestVersion, and an error, if there is any.
func (c *FakeLatestVersions) Update(ctx context.Context, latestVersion *v1alpha1.LatestVersion, opts v1.UpdateOptions) (result *v1alpha1.LatestVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(latestversionsResource, c.ns, latestVersion), &v1alpha1.LatestVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LatestVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLatestVersions) UpdateStatus(ctx context.Context, latestVersion *v1alpha1.LatestVersion, opts v1.UpdateOptions) (*v1alpha1.LatestVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(latestversionsResource, "status", c.ns, latestVersion), &v1alpha1.LatestVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LatestVersion), err
}

// Delete takes name of the latestVersion and deletes it. Returns an error if one occurs.
func (c *FakeLatestVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(latestversionsResource, c.ns, name, opts), &v1alpha1.LatestVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLatestVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(latestversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LatestVersionList{})
	return err
}

// Patch applies the patch and returns the patched latestVersion.
func (c *FakeLatestVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LatestVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(latestversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LatestVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LatestVersion), err
}