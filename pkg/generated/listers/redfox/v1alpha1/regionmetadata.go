// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/krafton-hq/redfox/pkg/apis/redfox/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RegionMetadataLister helps list RegionMetadatas.
// All objects returned here must be treated as read-only.
type RegionMetadataLister interface {
	// List lists all RegionMetadatas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegionMetadata, err error)
	// RegionMetadatas returns an object that can list and get RegionMetadatas.
	RegionMetadatas(namespace string) RegionMetadataNamespaceLister
	RegionMetadataListerExpansion
}

// regionMetadataLister implements the RegionMetadataLister interface.
type regionMetadataLister struct {
	indexer cache.Indexer
}

// NewRegionMetadataLister returns a new RegionMetadataLister.
func NewRegionMetadataLister(indexer cache.Indexer) RegionMetadataLister {
	return &regionMetadataLister{indexer: indexer}
}

// List lists all RegionMetadatas in the indexer.
func (s *regionMetadataLister) List(selector labels.Selector) (ret []*v1alpha1.RegionMetadata, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegionMetadata))
	})
	return ret, err
}

// RegionMetadatas returns an object that can list and get RegionMetadatas.
func (s *regionMetadataLister) RegionMetadatas(namespace string) RegionMetadataNamespaceLister {
	return regionMetadataNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegionMetadataNamespaceLister helps list and get RegionMetadatas.
// All objects returned here must be treated as read-only.
type RegionMetadataNamespaceLister interface {
	// List lists all RegionMetadatas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegionMetadata, err error)
	// Get retrieves the RegionMetadata from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RegionMetadata, error)
	RegionMetadataNamespaceListerExpansion
}

// regionMetadataNamespaceLister implements the RegionMetadataNamespaceLister
// interface.
type regionMetadataNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RegionMetadatas in the indexer for a given namespace.
func (s regionMetadataNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RegionMetadata, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegionMetadata))
	})
	return ret, err
}

// Get retrieves the RegionMetadata from the indexer for a given namespace and name.
func (s regionMetadataNamespaceLister) Get(name string) (*v1alpha1.RegionMetadata, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("regionmetadata"), name)
	}
	return obj.(*v1alpha1.RegionMetadata), nil
}
