// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/stolostron/cluster-lifecycle-api/clusterinfo/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedClusterInfos implements ManagedClusterInfoInterface
type FakeManagedClusterInfos struct {
	Fake *FakeInternalV1beta1
	ns   string
}

var managedclusterinfosResource = schema.GroupVersionResource{Group: "internal.open-cluster-management.io", Version: "v1beta1", Resource: "managedclusterinfos"}

var managedclusterinfosKind = schema.GroupVersionKind{Group: "internal.open-cluster-management.io", Version: "v1beta1", Kind: "ManagedClusterInfo"}

// Get takes name of the managedClusterInfo, and returns the corresponding managedClusterInfo object, and an error if there is any.
func (c *FakeManagedClusterInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ManagedClusterInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedclusterinfosResource, c.ns, name), &v1beta1.ManagedClusterInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ManagedClusterInfo), err
}

// List takes label and field selectors, and returns the list of ManagedClusterInfos that match those selectors.
func (c *FakeManagedClusterInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ManagedClusterInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedclusterinfosResource, managedclusterinfosKind, c.ns, opts), &v1beta1.ManagedClusterInfoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ManagedClusterInfoList{ListMeta: obj.(*v1beta1.ManagedClusterInfoList).ListMeta}
	for _, item := range obj.(*v1beta1.ManagedClusterInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedClusterInfos.
func (c *FakeManagedClusterInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedclusterinfosResource, c.ns, opts))

}

// Create takes the representation of a managedClusterInfo and creates it.  Returns the server's representation of the managedClusterInfo, and an error, if there is any.
func (c *FakeManagedClusterInfos) Create(ctx context.Context, managedClusterInfo *v1beta1.ManagedClusterInfo, opts v1.CreateOptions) (result *v1beta1.ManagedClusterInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedclusterinfosResource, c.ns, managedClusterInfo), &v1beta1.ManagedClusterInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ManagedClusterInfo), err
}

// Update takes the representation of a managedClusterInfo and updates it. Returns the server's representation of the managedClusterInfo, and an error, if there is any.
func (c *FakeManagedClusterInfos) Update(ctx context.Context, managedClusterInfo *v1beta1.ManagedClusterInfo, opts v1.UpdateOptions) (result *v1beta1.ManagedClusterInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedclusterinfosResource, c.ns, managedClusterInfo), &v1beta1.ManagedClusterInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ManagedClusterInfo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedClusterInfos) UpdateStatus(ctx context.Context, managedClusterInfo *v1beta1.ManagedClusterInfo, opts v1.UpdateOptions) (*v1beta1.ManagedClusterInfo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedclusterinfosResource, "status", c.ns, managedClusterInfo), &v1beta1.ManagedClusterInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ManagedClusterInfo), err
}

// Delete takes name of the managedClusterInfo and deletes it. Returns an error if one occurs.
func (c *FakeManagedClusterInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(managedclusterinfosResource, c.ns, name, opts), &v1beta1.ManagedClusterInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedClusterInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedclusterinfosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ManagedClusterInfoList{})
	return err
}

// Patch applies the patch and returns the patched managedClusterInfo.
func (c *FakeManagedClusterInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ManagedClusterInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedclusterinfosResource, c.ns, name, pt, data, subresources...), &v1beta1.ManagedClusterInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ManagedClusterInfo), err
}