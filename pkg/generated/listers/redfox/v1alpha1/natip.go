// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/krafton-hq/redfox/pkg/apis/redfox/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NatIpLister helps list NatIps.
// All objects returned here must be treated as read-only.
type NatIpLister interface {
	// List lists all NatIps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NatIp, err error)
	// NatIps returns an object that can list and get NatIps.
	NatIps(namespace string) NatIpNamespaceLister
	NatIpListerExpansion
}

// natIpLister implements the NatIpLister interface.
type natIpLister struct {
	indexer cache.Indexer
}

// NewNatIpLister returns a new NatIpLister.
func NewNatIpLister(indexer cache.Indexer) NatIpLister {
	return &natIpLister{indexer: indexer}
}

// List lists all NatIps in the indexer.
func (s *natIpLister) List(selector labels.Selector) (ret []*v1alpha1.NatIp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NatIp))
	})
	return ret, err
}

// NatIps returns an object that can list and get NatIps.
func (s *natIpLister) NatIps(namespace string) NatIpNamespaceLister {
	return natIpNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NatIpNamespaceLister helps list and get NatIps.
// All objects returned here must be treated as read-only.
type NatIpNamespaceLister interface {
	// List lists all NatIps in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NatIp, err error)
	// Get retrieves the NatIp from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NatIp, error)
	NatIpNamespaceListerExpansion
}

// natIpNamespaceLister implements the NatIpNamespaceLister
// interface.
type natIpNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NatIps in the indexer for a given namespace.
func (s natIpNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NatIp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NatIp))
	})
	return ret, err
}

// Get retrieves the NatIp from the indexer for a given namespace and name.
func (s natIpNamespaceLister) Get(name string) (*v1alpha1.NatIp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("natip"), name)
	}
	return obj.(*v1alpha1.NatIp), nil
}
