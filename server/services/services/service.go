package services

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type ClusterService[T any] interface {
	Init(ctx context.Context) error
	Get(ctx context.Context, name string) (T, error)
	List(ctx context.Context, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, obj T) error
	Update(ctx context.Context, obj T) error
	Delete(ctx context.Context, name string) error
}

type NamespacedService[T any] interface {
	Get(ctx context.Context, namespace string, name string) (T, error)
	List(ctx context.Context, labelSelectors map[string]string) ([]T, error)
	ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, obj T) error
	Update(ctx context.Context, obj T) error
	Delete(ctx context.Context, namespace string, name string) error
}

type NamespacedGvkService[T any] interface {
	Get(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, name string) (T, error)
	List(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, labelSelectors map[string]string) ([]T, error)
	ListNamespaced(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, obj T) error
	Update(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, obj T) error
	Delete(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, name string) error
}
