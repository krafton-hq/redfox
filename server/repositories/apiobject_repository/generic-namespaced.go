package apiobject_repository

import (
	"context"
	"fmt"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type GenericNamespacedRepository[T domain_helper.Metadatable] struct {
	repos map[string]ClusterRepository[T]
	gvk   *idl_common.GroupVersionKindSpec

	repoFactory ClusterRepositoryFactory[T]
}

func NewGenericNamespacedRepository[T domain_helper.Metadatable](gvk *idl_common.GroupVersionKindSpec, repoFactory ClusterRepositoryFactory[T]) *GenericNamespacedRepository[T] {
	return &GenericNamespacedRepository[T]{
		repos:       map[string]ClusterRepository[T]{},
		gvk:         gvk,
		repoFactory: repoFactory,
	}
}

func (r *GenericNamespacedRepository[T]) Get(ctx context.Context, namespace string, name string) (T, error) {
	repo, exist := r.repos[namespace]
	if !exist {
		var v T
		return v, errors.NewNotFound(fmt.Sprintf("Namespace/%s", namespace))
	}

	return repo.Get(ctx, name)
}

func (r *GenericNamespacedRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
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

func (r *GenericNamespacedRepository[T]) ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error) {
	repo, exist := r.repos[namespace]
	if !exist {
		return nil, errors.NewNotFound(fmt.Sprintf("Namespace/%s", namespace))
	}

	return repo.List(ctx, labelSelectors)
}

func (r *GenericNamespacedRepository[T]) Create(ctx context.Context, obj T) error {
	namespace := obj.GetMetadata().Namespace
	repo, exist := r.repos[namespace]
	if !exist {
		return errors.NewNotFound(fmt.Sprintf("Namespace/%s", namespace))
	}

	return repo.Create(ctx, obj)
}

func (r *GenericNamespacedRepository[T]) Update(ctx context.Context, obj T) error {
	namespace := obj.GetMetadata().Namespace
	repo, exist := r.repos[namespace]
	if !exist {
		return errors.NewNotFound(fmt.Sprintf("Namespace/%s", namespace))
	}

	return repo.Update(ctx, obj)
}

func (r *GenericNamespacedRepository[T]) Delete(ctx context.Context, namespace string, name string) error {
	repo, exist := r.repos[namespace]
	if !exist {
		return errors.NewNotFound(fmt.Sprintf("Namespace/%s", namespace))
	}

	return repo.Delete(ctx, name)
}

func (r *GenericNamespacedRepository[T]) Truncate(ctx context.Context) error {
	for _, repo := range r.repos {
		if err := repo.Truncate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r *GenericNamespacedRepository[T]) Info() *idl_common.GroupVersionKindSpec {
	return proto.Clone(r.gvk).(*idl_common.GroupVersionKindSpec)
}

func (r *GenericNamespacedRepository[T]) EnableNamespace(ctx context.Context, namespace string) bool {
	if _, exist := r.repos[namespace]; exist {
		return false
	}
	r.repos[namespace] = r.repoFactory.Create(r.gvk, namespace)
	return true
}

func (r *GenericNamespacedRepository[T]) DisableNamespace(ctx context.Context, namespace string) bool {
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
