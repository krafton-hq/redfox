package apiobject_repository

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/namespaces"
	"google.golang.org/protobuf/proto"
)

type SimpleNamespacedRepository[T Metadatable] struct {
	repos map[string]ClusterRepository[T]
	gvk   *namespaces.GroupVersionKind
}

func (r *SimpleNamespacedRepository[T]) Get(ctx context.Context, namespace string, name string) (T, error) {
	repo, exist := r.repos[namespace]
	if !exist {
		var v T
		return v, fmt.Errorf("no Namespace/%s resource found  in cluster", namespace)
	}

	return repo.Get(ctx, name)
}

func (r *SimpleNamespacedRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
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

func (r *SimpleNamespacedRepository[T]) ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error) {
	repo, exist := r.repos[namespace]
	if !exist {
		return nil, fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.List(ctx, labelSelectors)
}

func (r *SimpleNamespacedRepository[T]) Create(ctx context.Context, obj T) error {
	namespace := obj.GetMetadata().Namespace
	repo, exist := r.repos[namespace]
	if !exist {
		return fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.Create(ctx, obj)
}

func (r *SimpleNamespacedRepository[T]) Update(ctx context.Context, obj T) error {
	namespace := obj.GetMetadata().Namespace
	repo, exist := r.repos[namespace]
	if !exist {
		return fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.Update(ctx, obj)
}

func (r *SimpleNamespacedRepository[T]) Delete(ctx context.Context, namespace string, name string) error {
	repo, exist := r.repos[namespace]
	if !exist {
		return fmt.Errorf("requested namespace is not eixst: %s", namespace)
	}

	return repo.Delete(ctx, name)
}

func (r *SimpleNamespacedRepository[T]) Info() *namespaces.GroupVersionKind {
	return proto.Clone(r.gvk).(*namespaces.GroupVersionKind)
}

func (r *SimpleNamespacedRepository[T]) EnableNamespace(ctx context.Context, namespace string) bool {
	if _, exist := r.repos[namespace]; exist {
		return false
	}
	r.repos[namespace] = NewInMemoryClusterRepositoryWithFixedNamespace[T](r.gvk, namespace)
	return true
}

func (r *SimpleNamespacedRepository[T]) DisableNamespace(ctx context.Context, namespace string) bool {
	if _, exist := r.repos[namespace]; !exist {
		return false
	}
	delete(r.repos, namespace)
	return true
}
