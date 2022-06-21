package transactional

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type Transactional interface {
	WithDatabaseContext(ctx context.Context) context.Context
	WithTransaction(ctx context.Context, fn func(context.Context, *sqlx.Tx, goqu.SQLDialect) error) error
}
