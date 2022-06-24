package transactional

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"go.uber.org/zap"
)

type transactionKey struct{}
type databaseKey struct{}

const DialectMysql = "mysql"

type sqlt struct {
	db       *sqlx.DB
	dialect  goqu.SQLDialect
	txOption *sql.TxOptions
}

type SqlOption struct {
	txOption *sql.TxOptions
}

// NewSqlTransactional : SqlOption parameter is nullable
func NewSqlTransactional(db *sqlx.DB, dialectName string, option *SqlOption) (*sqlt, error) {
	if db == nil {
		return nil, errors.NewInvalidArguments("'db' should be not null")
	}

	dialect := goqu.GetDialect(dialectName)
	if dialect.Dialect() != dialectName {
		return nil, errors.NewErrorf("Failed to Get Dialect expect: %s, actual: %s", dialectName, dialect.Dialect())
	}

	var txOption *sql.TxOptions
	if option != nil {
		txOption = option.txOption
	}

	return &sqlt{db: db, dialect: dialect, txOption: txOption}, nil
}

func (t *sqlt) WithDatabaseContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, databaseKey{}, t.db)
}

func (t *sqlt) WithTransaction(ctx context.Context, fn func(context.Context, *sqlx.Tx, goqu.SQLDialect) error) error {
	if tx, found := ctx.Value(transactionKey{}).(*sqlx.Tx); found {
		return fn(ctx, tx, t.dialect)
	}

	db, found := ctx.Value(databaseKey{}).(*sqlx.DB)
	if !found {
		message := "Sqlt.WithTransaction: Cannot find Database object in current context.Context"
		zap.S().Info(message)
		return errors.NewInternalError(message)
	}
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		message := "Sqlt.WithTransaction: Failed to Start Transaction"
		zap.S().Infow(message, "error", err)
		return errors.WrapInternalError(err, message)
	}

	defer func() {
		if err == nil {
			return
		}
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = errors.WrapInternalError(err, fmt.Sprintf("Sqlt.WithTransaction: Failed to rollback %s", rollbackErr.Error()))
		}
	}()

	ctx = context.WithValue(ctx, transactionKey{}, tx)
	err = fn(ctx, tx, t.dialect)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = errors.WrapInternalError(err, "Sqlt.WithTransaction: Failed to Commit Transaction")
		return err
	}
	return err
}
