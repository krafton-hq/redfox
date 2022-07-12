package service_decorator

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/transactional"
	"github.com/krafton-hq/red-fox/server/services/services"
)

type TransactionalNamespacedGvkService[T any] struct {
	service services.NamespacedGvkService[T]

	tr transactional.Transactional
}

func NewTransactionalNamespacedGvkService[T any](service services.NamespacedGvkService[T], tr transactional.Transactional) services.NamespacedGvkService[T] {
	return &TransactionalNamespacedGvkService[T]{service: service, tr: tr}
}

func (s *TransactionalNamespacedGvkService[T]) Get(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, name string) (T, error) {
	ctx = s.tr.WithDatabaseContext(ctx)

	var result T
	var err error
	err = s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		result, err = s.service.Get(ctx, gvk, namespace, name)
		return err
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *TransactionalNamespacedGvkService[T]) List(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, labelSelectors map[string]string) ([]T, error) {
	ctx = s.tr.WithDatabaseContext(ctx)

	var result []T
	var err error
	err = s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		result, err = s.service.List(ctx, gvk, labelSelectors)
		return err
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *TransactionalNamespacedGvkService[T]) ListNamespaced(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, labelSelectors map[string]string) ([]T, error) {
	ctx = s.tr.WithDatabaseContext(ctx)

	var result []T
	var err error
	err = s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		result, err = s.service.ListNamespaced(ctx, gvk, namespace, labelSelectors)
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TransactionalNamespacedGvkService[T]) Create(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, obj T) error {
	ctx = s.tr.WithDatabaseContext(ctx)

	err := s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		return s.service.Create(ctx, gvk, obj)
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionalNamespacedGvkService[T]) Update(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, obj T) error {
	ctx = s.tr.WithDatabaseContext(ctx)

	err := s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		return s.service.Update(ctx, gvk, obj)
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionalNamespacedGvkService[T]) Delete(ctx context.Context, gvk *idl_common.GroupVersionKindSpec, namespace string, name string) error {
	ctx = s.tr.WithDatabaseContext(ctx)

	err := s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		return s.service.Delete(ctx, gvk, namespace, name)
	})
	if err != nil {
		return err
	}
	return nil
}
