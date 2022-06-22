package apiobject_repository

import (
	"context"

	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
)

const DefaultSystemNamespace = "@fox-system"

type ClusterRepository[T domain_helper.Metadatable] interface {
	Get(ctx context.Context, name string) (T, error)
	List(ctx context.Context, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, obj T) error
	Update(ctx context.Context, obj T) error
	Delete(ctx context.Context, name string) error
	Truncate(ctx context.Context) error

	Info() *namespaces.GroupVersionKind
}

type NamespacedRepository[T domain_helper.Metadatable] interface {
	Get(ctx context.Context, namespace string, name string) (T, error)
	List(ctx context.Context, labelSelectors map[string]string) ([]T, error)
	ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error)
	Create(ctx context.Context, obj T) error
	Update(ctx context.Context, obj T) error
	Delete(ctx context.Context, namespace string, name string) error
	Truncate(ctx context.Context) error

	NamespacedRepositoryMetadata
}

type NamespacedRepositoryMetadata interface {
	Info() *namespaces.GroupVersionKind
	EnableNamespace(ctx context.Context, namespace string) bool
	DisableNamespace(ctx context.Context, namespace string) bool
}

type ClusterRepositoryFactory[T domain_helper.Metadatable] interface {
	Create(gvk *namespaces.GroupVersionKind, uniqueKeySuffix string) ClusterRepository[T]
}
