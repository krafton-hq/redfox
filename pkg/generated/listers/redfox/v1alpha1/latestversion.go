// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/krafton-hq/redfox/pkg/apis/redfox/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LatestVersionLister helps list LatestVersions.
// All objects returned here must be treated as read-only.
type LatestVersionLister interface {
	// List lists all LatestVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LatestVersion, err error)
	// LatestVersions returns an object that can list and get LatestVersions.
	LatestVersions(namespace string) LatestVersionNamespaceLister
	LatestVersionListerExpansion
}

// latestVersionLister implements the LatestVersionLister interface.
type latestVersionLister struct {
	indexer cache.Indexer
}

// NewLatestVersionLister returns a new LatestVersionLister.
func NewLatestVersionLister(indexer cache.Indexer) LatestVersionLister {
	return &latestVersionLister{indexer: indexer}
}

// List lists all LatestVersions in the indexer.
func (s *latestVersionLister) List(selector labels.Selector) (ret []*v1alpha1.LatestVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LatestVersion))
	})
	return ret, err
}

// LatestVersions returns an object that can list and get LatestVersions.
func (s *latestVersionLister) LatestVersions(namespace string) LatestVersionNamespaceLister {
	return latestVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LatestVersionNamespaceLister helps list and get LatestVersions.
// All objects returned here must be treated as read-only.
type LatestVersionNamespaceLister interface {
	// List lists all LatestVersions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LatestVersion, err error)
	// Get retrieves the LatestVersion from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LatestVersion, error)
	LatestVersionNamespaceListerExpansion
}

// latestVersionNamespaceLister implements the LatestVersionNamespaceLister
// interface.
type latestVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LatestVersions in the indexer for a given namespace.
func (s latestVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LatestVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LatestVersion))
	})
	return ret, err
}

// Get retrieves the LatestVersion from the indexer for a given namespace and name.
func (s latestVersionNamespaceLister) Get(name string) (*v1alpha1.LatestVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("latestversion"), name)
	}
	return obj.(*v1alpha1.LatestVersion), nil
}
