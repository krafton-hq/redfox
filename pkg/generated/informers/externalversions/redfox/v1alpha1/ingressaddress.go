// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	redfoxv1alpha1 "github.com/krafton-hq/redfox/pkg/apis/redfox/v1alpha1"
	versioned "github.com/krafton-hq/redfox/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/krafton-hq/redfox/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/krafton-hq/redfox/pkg/generated/listers/redfox/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IngressAddressInformer provides access to a shared informer and lister for
// IngressAddresses.
type IngressAddressInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IngressAddressLister
}

type ingressAddressInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIngressAddressInformer constructs a new informer for IngressAddress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressAddressInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIngressAddressInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIngressAddressInformer constructs a new informer for IngressAddress type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressAddressInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetadataV1alpha1().IngressAddresses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetadataV1alpha1().IngressAddresses(namespace).Watch(context.TODO(), options)
			},
		},
		&redfoxv1alpha1.IngressAddress{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressAddressInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIngressAddressInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ingressAddressInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&redfoxv1alpha1.IngressAddress{}, f.defaultInformer)
}

func (f *ingressAddressInformer) Lister() v1alpha1.IngressAddressLister {
	return v1alpha1.NewIngressAddressLister(f.Informer().GetIndexer())
}
