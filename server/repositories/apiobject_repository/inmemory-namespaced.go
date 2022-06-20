package apiobject_repository

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type InmemoryNamespacedRepository[T domain_helper.Metadatable] struct {
	repos map[string]ClusterRepository[T]
	gvk   *namespaces.GroupVersionKind

	repoFactory ClusterRepositoryFactory[T]
}

func NewInmemoryNamespacedRepository[T domain_helper.Metadatable](gvk *namespaces.GroupVersionKind, repoFactory ClusterRepositoryFactory[T]) *InmemoryNamespacedRepository[T] {
	return &InmemoryNamespacedRepository[T]{
		repos:       map[string]ClusterRepository[T]{},
		gvk:         gvk,
		repoFactory: repoFactory,
	}
}

func (r *InmemoryNamespacedRepository[T]) Get(ctx context.Context, namespace string, name string) (T, error) {
	repo, exist := r.repos[namespace]
	if !exist {
		var v T
		return v, fmt.Errorf("no Namespace/%s resource found  in cluster", namespace)
	}

	return repo.Get(ctx, name)
}

func (r *InmemoryNamespacedRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	var ret []T
	for _, repo := range r.repos {
		nRet, err := repo.List(ctx, labelSelectors)
		if err != nil {
			return nil, err
		}
		ret = append(ret, nRet...)
	}
	return ret, nil
}

func (r *InmemoryNamespacedRepository[T]) ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error) {
	repo, exist := r.repos[namespace]
	if !exist {
		return nil, fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.List(ctx, labelSelectors)
}

func (r *InmemoryNamespacedRepository[T]) Create(ctx context.Context, obj T) error {
	namespace := obj.GetMetadata().Namespace
	repo, exist := r.repos[namespace]
	if !exist {
		return fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.Create(ctx, obj)
}

func (r *InmemoryNamespacedRepository[T]) Update(ctx context.Context, obj T) error {
	namespace := obj.GetMetadata().Namespace
	repo, exist := r.repos[namespace]
	if !exist {
		return fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.Update(ctx, obj)
}

func (r *InmemoryNamespacedRepository[T]) Delete(ctx context.Context, namespace string, name string) error {
	repo, exist := r.repos[namespace]
	if !exist {
		return fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.Delete(ctx, name)
}

func (r *InmemoryNamespacedRepository[T]) Info() *namespaces.GroupVersionKind {
	return proto.Clone(r.gvk).(*namespaces.GroupVersionKind)
}

func (r *InmemoryNamespacedRepository[T]) EnableNamespace(ctx context.Context, namespace string) bool {
	if _, exist := r.repos[namespace]; exist {
		return false
	}
	r.repos[namespace] = r.repoFactory.Create(r.gvk, namespace)
	return true
}

func (r *InmemoryNamespacedRepository[T]) DisableNamespace(ctx context.Context, namespace string) bool {
	repo, exist := r.repos[namespace]
	if !exist {
		return false
	}
	err := repo.Truncate(ctx)
	if err != nil {
		zap.S().Infow("Delete Namespace failed", "error", err)
		return false
	}
	delete(r.repos, namespace)
	return true
}
