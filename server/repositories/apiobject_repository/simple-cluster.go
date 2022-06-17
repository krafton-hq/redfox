package apiobject_repository

import (
	"context"
	"fmt"
)

const defaultSystemNamespace = "@fox-system"

type SimpleClusterRepository[T Metadatable] struct {
	nsRepository    NamespacedRepository[T]
	systemNamespace string
}

func NewSimpleClusterRepository[T Metadatable](nsRepository NamespacedRepository[T]) *SimpleClusterRepository[T] {
	return NewSimpleClusterRepository2[T](nsRepository, defaultSystemNamespace)
}

func NewSimpleClusterRepository2[T Metadatable](nsRepository NamespacedRepository[T], fixedNamespace string) *SimpleClusterRepository[T] {
	nsRepository.EnableNamespace(context.TODO(), fixedNamespace)

	return &SimpleClusterRepository[T]{
		nsRepository:    nsRepository,
		systemNamespace: fixedNamespace,
	}
}

func (r *SimpleClusterRepository[T]) Get(ctx context.Context, name string) (T, error) {
	obj, err := r.nsRepository.Get(ctx, r.systemNamespace, name)
	if err != nil {
		var v T
		return v, err
	}
	obj.GetMetadata().Namespace = ""
	return obj, nil
}

func (r *SimpleClusterRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	objects, err := r.nsRepository.ListNamespaced(ctx, r.systemNamespace, labelSelectors)
	if err != nil {
		return nil, err
	}
	for _, obj := range objects {
		obj.GetMetadata().Namespace = ""
	}
	return objects, nil
}

func (r *SimpleClusterRepository[T]) Create(ctx context.Context, obj T) error {
	if obj.GetMetadata().Namespace != "" {
		return fmt.Errorf("ApiObject should not have namespace")
	}

	obj.GetMetadata().Namespace = r.systemNamespace
	return r.nsRepository.Create(ctx, obj)
}

func (r *SimpleClusterRepository[T]) Update(ctx context.Context, obj T) error {
	if obj.GetMetadata().Namespace != "" {
		return fmt.Errorf("ApiObject should not have namespace")
	}

	obj.GetMetadata().Namespace = r.systemNamespace
	return r.nsRepository.Update(ctx, obj)
}

func (r *SimpleClusterRepository[T]) Delete(ctx context.Context, name string) error {
	return r.nsRepository.Delete(ctx, r.systemNamespace, name)
}
