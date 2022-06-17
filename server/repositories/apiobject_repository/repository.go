package apiobject_repository

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
)

type Metadatable interface {
	GetMetadata() *idl_common.ObjectMeta
}

type ClusterRepository[T Metadatable] interface {
	Get(ctx context.Context, name string) (T, error)
	List(ctx context.Context, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, obj T) error
	Update(ctx context.Context, obj T) error
	Delete(ctx context.Context, name string) error

	Info() *namespaces.GroupVersionKind
}

type NamespacedRepository[T Metadatable] interface {
	Get(ctx context.Context, namespace string, name string) (T, error)
	List(ctx context.Context, labelSelectors map[string]string) ([]T, error)
	ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, obj T) error
	Update(ctx context.Context, obj T) error
	Delete(ctx context.Context, namespace string, name string) error

	NamespacedRepositoryMetadata
}

type NamespacedRepositoryMetadata interface {
	Info() *namespaces.GroupVersionKind
	EnableNamespace(ctx context.Context, namespace string) bool
	DisableNamespace(ctx context.Context, namespace string) bool
}
