package apiobject_repository

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"google.golang.org/protobuf/proto"
)

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

type InMemoryClusterRepository[T domain_helper.Metadatable] struct {
	local map[string]T
	gvk   *namespaces.GroupVersionKind

	uniqueKey string
}

func NewInMemoryClusterRepository[T domain_helper.Metadatable](gvk *namespaces.GroupVersionKind, uniqueKeySuffix string) *InMemoryClusterRepository[T] {
	return &InMemoryClusterRepository[T]{
		local:     map[string]T{},
		gvk:       proto.Clone(gvk).(*namespaces.GroupVersionKind),
		uniqueKey: fmt.Sprintf("%s/%s/%s", gvk.Group, gvk.Kind, uniqueKeySuffix),
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
	if _, exist := r.local[name]; exist {
		return fmt.Errorf("object '%s' alreay eixst", name)
	}
	r.local[name] = obj
	return nil
}

func (r *InMemoryClusterRepository[T]) Update(ctx context.Context, obj T) error {
	name := obj.GetMetadata().Name
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

func (r *InMemoryClusterRepository[T]) Truncate(ctx context.Context) error {
	r.local = map[string]T{}
	return nil
}

func (r *InMemoryClusterRepository[T]) Info() *namespaces.GroupVersionKind {
	return proto.Clone(r.gvk).(*namespaces.GroupVersionKind)
}

type InmemoryClusterRepositoryFactory[T domain_helper.Metadatable] struct {
}

func (f *InmemoryClusterRepositoryFactory[T]) Create(gvk *namespaces.GroupVersionKind, uniqueKeySuffix string) ClusterRepository[T] {
	return NewInMemoryClusterRepository[T](gvk, uniqueKeySuffix)
}

func NewInmemoryClusterRepositoryFactory[T domain_helper.Metadatable]() *InmemoryClusterRepositoryFactory[T] {
	return &InmemoryClusterRepositoryFactory[T]{}
}
