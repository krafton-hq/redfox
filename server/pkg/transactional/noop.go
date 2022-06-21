package transactional

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type noop struct {
}

func NewNoop() *noop {
	return &noop{}
}

func (n *noop) WithDatabaseContext(ctx context.Context) context.Context {
	return ctx
}

func (n *noop) WithTransaction(ctx context.Context, fn func(context.Context, *sqlx.Tx, goqu.SQLDialect) error) error {
	return fn(ctx, nil, nil)
}
