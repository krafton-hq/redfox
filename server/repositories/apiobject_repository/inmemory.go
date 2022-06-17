package apiobject_repository

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"google.golang.org/protobuf/proto"
)

type InMemoryNamespacedRepository[T Metadatable] struct {
	local map[string]map[string]T
	gvk   *namespaces.GroupVersionKind
}

func NewInMemoryNamespacedRepository[T Metadatable](gvk *namespaces.GroupVersionKind) *InMemoryNamespacedRepository[T] {
	return &InMemoryNamespacedRepository[T]{
		local: map[string]map[string]T{},
		gvk:   gvk,
	}
}

func (r *InMemoryNamespacedRepository[T]) Get(ctx context.Context, namespace string, name string) (T, error) {
	ns, exist := r.local[namespace]
	if !exist {
		var v T
		return v, fmt.Errorf("namespace not eixst")
	}

	if nat, exist := ns[name]; exist {
		return nat, nil
	} else {
		var v T
		return v, fmt.Errorf("natip not eixst")
	}
}

func (r *InMemoryNamespacedRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	var ret []T
	for _, namespaced := range r.local {
		for _, objects := range namespaced {
			if containsLabels(objects.GetMetadata(), labelSelectors) {
				ret = append(ret, objects)
			}
		}
	}
	return ret, nil
}

func (r *InMemoryNamespacedRepository[T]) ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error) {
	ns, exist := r.local[namespace]
	if !exist {
		return nil, fmt.Errorf("namespace not eixst")
	}

	var ret []T
	for _, object := range ns {
		if containsLabels(object.GetMetadata(), labelSelectors) {
			ret = append(ret, object)
		}
	}
	return ret, nil
}

func containsLabels(metadata *idl_common.ObjectMeta, labels map[string]string) bool {
	for key, value := range labels {
		if !containsLabel(metadata, key, value) {
			return false
		}
	}
	return true
}

func containsLabel(metadata *idl_common.ObjectMeta, tkey, tvalue string) bool {
	for lkey, lvalue := range metadata.Labels {
		if lkey == tkey {
			if tvalue == "" {
				return true
			}
			if lvalue == tvalue {
				return true
			}
		}
	}
	return false
}

func (r *InMemoryNamespacedRepository[T]) Create(ctx context.Context, obj T) error {
	ns, exist := r.local[obj.GetMetadata().Namespace]
	if !exist {
		return fmt.Errorf("namespace not eixst")
	}

	name := obj.GetMetadata().Name
	if _, exist := ns[name]; exist {
		return fmt.Errorf("object '%s' alreay eixst", name)
	}
	ns[name] = obj
	return nil
}

func (r *InMemoryNamespacedRepository[T]) Update(ctx context.Context, obj T) error {
	ns, exist := r.local[obj.GetMetadata().Namespace]
	if !exist {
		return fmt.Errorf("namespace not eixst")
	}

	name := obj.GetMetadata().Name
	if _, exist := ns[name]; !exist {
		return fmt.Errorf("natip not eixst")
	}
	ns[name] = obj
	return nil
}

func (r *InMemoryNamespacedRepository[T]) Delete(ctx context.Context, namespace string, name string) error {
	ns, exist := r.local[namespace]
	if !exist {
		return fmt.Errorf("namespace not eixst")
	}

	if _, exist := ns[name]; !exist {
		return fmt.Errorf("natip not eixst")
	}
	delete(ns, name)
	return nil
}

func (r *InMemoryNamespacedRepository[T]) EnableNamespace(ctx context.Context, namespace string) bool {
	if _, exist := r.local[namespace]; exist {
		return false
	}
	r.local[namespace] = map[string]T{}
	return true
}

func (r *InMemoryNamespacedRepository[T]) DisableNamespace(ctx context.Context, namespace string) bool {
	if _, exist := r.local[namespace]; !exist {
		return false
	}
	delete(r.local, namespace)
	return true
}

func (r *InMemoryNamespacedRepository[T]) Info() *namespaces.GroupVersionKind {
	return proto.Clone(r.gvk).(*namespaces.GroupVersionKind)
}

type InMemoryClusterRepository[T Metadatable] struct {
	local map[string]T
	gvk   *namespaces.GroupVersionKind

	isNamespaced   bool
	fixedNamespace string
}

func NewInMemoryClusterRepository[T Metadatable](gvk *namespaces.GroupVersionKind) *InMemoryClusterRepository[T] {
	return &InMemoryClusterRepository[T]{
		local:          map[string]T{},
		gvk:            proto.Clone(gvk).(*namespaces.GroupVersionKind),
		isNamespaced:   false,
		fixedNamespace: defaultSystemNamespace,
	}
}

func NewInMemoryClusterRepositoryWithFixedNamespace[T Metadatable](gvk *namespaces.GroupVersionKind, fixedNamespace string) *InMemoryClusterRepository[T] {
	return &InMemoryClusterRepository[T]{
		local:          map[string]T{},
		gvk:            proto.Clone(gvk).(*namespaces.GroupVersionKind),
		isNamespaced:   true,
		fixedNamespace: fixedNamespace,
	}
}

func (r *InMemoryClusterRepository[T]) Get(ctx context.Context, name string) (T, error) {
	if obj, exist := r.local[name]; exist {
		return obj, nil
	}
	var v T
	return v, fmt.Errorf("cannot find '%s' object", name)
}

func (r *InMemoryClusterRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	var ret []T
	for _, object := range r.local {
		if containsLabels(object.GetMetadata(), labelSelectors) {
			ret = append(ret, object)
		}
	}
	return ret, nil
}

func (r *InMemoryClusterRepository[T]) Create(ctx context.Context, obj T) error {
	name := obj.GetMetadata().Name
	if r.isNamespaced && r.fixedNamespace != obj.GetMetadata().Namespace {
		return fmt.Errorf("invalid Namespaced Object ns: %s, name: %s", obj.GetMetadata().Namespace, name)
	}

	if _, exist := r.local[name]; exist {
		return fmt.Errorf("object '%s' alreay eixst", name)
	}
	r.local[name] = obj
	return nil
}

func (r *InMemoryClusterRepository[T]) Update(ctx context.Context, obj T) error {
	name := obj.GetMetadata().Name
	if r.isNamespaced && r.fixedNamespace != obj.GetMetadata().Namespace {
		return fmt.Errorf("invalid Namespaced Object ns: %s, name: %s", obj.GetMetadata().Namespace, name)
	}

	if _, exist := r.local[name]; !exist {
		return fmt.Errorf("object '%s' not eixst", name)
	}
	r.local[name] = obj
	return nil
}

func (r *InMemoryClusterRepository[T]) Delete(ctx context.Context, name string) error {
	if _, exist := r.local[name]; !exist {
		return fmt.Errorf("object '%s' not eixst", name)
	}
	delete(r.local, name)
	return nil
}

func (r *InMemoryClusterRepository[T]) Info() *namespaces.GroupVersionKind {
	return proto.Clone(r.gvk).(*namespaces.GroupVersionKind)
}
