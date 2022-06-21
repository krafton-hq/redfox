package service_helper

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/server/pkg/transactional"
)

type TransactionalNamespacedService[T any] struct {
	service NamespacedService[T]

	tr transactional.Transactional
}

func NewTransactionalNamespacedService[T any](service NamespacedService[T], tr transactional.Transactional) NamespacedService[T] {
	return &TransactionalNamespacedService[T]{service: service, tr: tr}
}

func (s *TransactionalNamespacedService[T]) Get(ctx context.Context, namespace string, name string) (T, error) {
	ctx = s.tr.WithDatabaseContext(ctx)

	var result T
	var err error
	err = s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		result, err = s.service.Get(ctx, namespace, name)
		return err
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *TransactionalNamespacedService[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	ctx = s.tr.WithDatabaseContext(ctx)

	var result []T
	var err error
	err = s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		result, err = s.service.List(ctx, labelSelectors)
		return err
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *TransactionalNamespacedService[T]) ListNamespaced(ctx context.Context, namespace string, labelSelectors map[string]string) ([]T, error) {
	ctx = s.tr.WithDatabaseContext(ctx)

	var result []T
	var err error
	err = s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		result, err = s.service.ListNamespaced(ctx, namespace, labelSelectors)
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TransactionalNamespacedService[T]) Create(ctx context.Context, obj T) error {
	ctx = s.tr.WithDatabaseContext(ctx)

	err := s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		return s.service.Create(ctx, obj)
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionalNamespacedService[T]) Update(ctx context.Context, obj T) error {
	ctx = s.tr.WithDatabaseContext(ctx)

	err := s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		return s.service.Update(ctx, obj)
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionalNamespacedService[T]) Delete(ctx context.Context, namespace string, name string) error {
	ctx = s.tr.WithDatabaseContext(ctx)

	err := s.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		return s.service.Delete(ctx, namespace, name)
	})
	if err != nil {
		return err
	}
	return nil
}
