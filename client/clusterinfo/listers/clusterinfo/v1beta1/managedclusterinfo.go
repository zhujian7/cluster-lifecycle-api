// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/stolostron/cluster-lifecycle-api/clusterinfo/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedClusterInfoLister helps list ManagedClusterInfos.
// All objects returned here must be treated as read-only.
type ManagedClusterInfoLister interface {
	// List lists all ManagedClusterInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ManagedClusterInfo, err error)
	// ManagedClusterInfos returns an object that can list and get ManagedClusterInfos.
	ManagedClusterInfos(namespace string) ManagedClusterInfoNamespaceLister
	ManagedClusterInfoListerExpansion
}

// managedClusterInfoLister implements the ManagedClusterInfoLister interface.
type managedClusterInfoLister struct {
	indexer cache.Indexer
}

// NewManagedClusterInfoLister returns a new ManagedClusterInfoLister.
func NewManagedClusterInfoLister(indexer cache.Indexer) ManagedClusterInfoLister {
	return &managedClusterInfoLister{indexer: indexer}
}

// List lists all ManagedClusterInfos in the indexer.
func (s *managedClusterInfoLister) List(selector labels.Selector) (ret []*v1beta1.ManagedClusterInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ManagedClusterInfo))
	})
	return ret, err
}

// ManagedClusterInfos returns an object that can list and get ManagedClusterInfos.
func (s *managedClusterInfoLister) ManagedClusterInfos(namespace string) ManagedClusterInfoNamespaceLister {
	return managedClusterInfoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedClusterInfoNamespaceLister helps list and get ManagedClusterInfos.
// All objects returned here must be treated as read-only.
type ManagedClusterInfoNamespaceLister interface {
	// List lists all ManagedClusterInfos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ManagedClusterInfo, err error)
	// Get retrieves the ManagedClusterInfo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ManagedClusterInfo, error)
	ManagedClusterInfoNamespaceListerExpansion
}

// managedClusterInfoNamespaceLister implements the ManagedClusterInfoNamespaceLister
// interface.
type managedClusterInfoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedClusterInfos in the indexer for a given namespace.
func (s managedClusterInfoNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ManagedClusterInfo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ManagedClusterInfo))
	})
	return ret, err
}

// Get retrieves the ManagedClusterInfo from the indexer for a given namespace and name.
func (s managedClusterInfoNamespaceLister) Get(name string) (*v1beta1.ManagedClusterInfo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("managedclusterinfo"), name)
	}
	return obj.(*v1beta1.ManagedClusterInfo), nil
}
