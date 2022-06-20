package apiobject_repository

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/json"
)

type transactionKey struct{}
type databaseKey struct{}

const tableApiObject = "api_object"
const fieldName = "name"
const fieldRepoKey = "repo_key"
const dialectMysql = "mysql"

func withTransaction(ctx context.Context, fn func(context.Context, *sqlx.Tx, goqu.SQLDialect) error) error {
	if tx, exist := ctx.Value(transactionKey{}).(*sqlx.Tx); exist {
		return fn(ctx, tx, goqu.GetDialect(dialectMysql))
	}

	db, exist := ctx.Value(databaseKey{}).(*sqlx.DB)
	if !exist {
		return fmt.Errorf("mysql.withTransaction: cannot find Database object in current context.Context")
	}
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("mysql.withTransaction: failed to start transaction %w", err)
	}
	ctx = context.WithValue(ctx, transactionKey{}, tx)

	defer func() {
		if err == nil {
			return
		}
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = fmt.Errorf("mysql.withTransaction: failed to rollback (%s) %w", rollbackErr.Error(), err)
		}
	}()

	err = fn(ctx, tx, goqu.GetDialect(dialectMysql))
	if err != nil {
		err = fmt.Errorf("mysql.withTransaction: failed to execute transaction %w", err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("postgres.WithCommit: failed to commit transaction %w", err)
		return err
	}
	return err
}

type ApiObject struct {
	Uid        int    `db:"uid" goqu:"skipinsert,skipupdate"`
	RepoKey    string `db:"repo_key" goqu:"skipupdate"`
	Name       string `db:"name" goqu:"skipupdate"`
	ObjectJson string `db:"object_json"`
	Labels     Labels `db:"labels" goqu:"json"`
}

type Labels []string

func (l *Labels) Scan(v interface{}) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, l)
	case string:
		return json.Unmarshal([]byte(vv), l)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func fromDbApiObject[T domain_helper.Metadatable](repo *MysqlClusterRepository[T], obj *ApiObject) (T, error) {
	t := repo.factory.Create()
	err := json.Unmarshal([]byte(obj.ObjectJson), t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func toApiObject[T domain_helper.Metadatable](repo *MysqlClusterRepository[T], t T) (*ApiObject, error) {
	apiObject := &ApiObject{}
	apiObject.Name = t.GetMetadata().Name
	apiObject.RepoKey = repo.uniqueKey
	apiObject.Labels = lo.Keys[string, string](t.GetMetadata().GetLabels())

	object, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	apiObject.ObjectJson = string(object)
	return apiObject, nil
}

type MysqlClusterRepository[T domain_helper.Metadatable] struct {
	gvk       *namespaces.GroupVersionKind
	uniqueKey string
	factory   domain_helper.MetadatableFactory[T]
}

func NewMysqlClusterRepository[T domain_helper.Metadatable](gvk *namespaces.GroupVersionKind, uniqueKeySuffix string, factory domain_helper.MetadatableFactory[T]) *MysqlClusterRepository[T] {
	return &MysqlClusterRepository[T]{
		gvk:       proto.Clone(gvk).(*namespaces.GroupVersionKind),
		uniqueKey: fmt.Sprintf("%s/%s/%s", gvk.Group, gvk.Kind, uniqueKeySuffix),
		factory:   factory,
	}
}

//goqu.L("? member of ")
func (r *MysqlClusterRepository[T]) Get(ctx context.Context, name string) (T, error) {
	dbApiObject := &ApiObject{}

	err := withTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Select("*").SetDialect(dialect).From(tableApiObject).Where(goqu.And(goqu.C(fieldName).Eq(name), goqu.C(fieldRepoKey).Eq(r.uniqueKey))).ToSQL()
		row := tx.QueryRowx(query)

		return row.StructScan(dbApiObject)
	})
	if err != nil {
		var t T
		return t, err
	}

	return fromDbApiObject[T](r, dbApiObject)
}

func (r *MysqlClusterRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	var dbApiObjects []*ApiObject
	err := withTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Select("*").SetDialect(dialect).From(tableApiObject).ToSQL()
		return tx.Select(dbApiObjects, query)
	})
	if err != nil {
		return nil, err
	}

	ts := lo.Map[*ApiObject, T](dbApiObjects, func(apiObject *ApiObject, _ int) T {
		var t T
		t, err = fromDbApiObject[T](r, apiObject)
		return t
	})
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *MysqlClusterRepository[T]) Create(ctx context.Context, obj T) error {
	dbApiObject, err := toApiObject[T](r, obj)
	if err != nil {
		return err
	}
	err = withTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Insert(tableApiObject).SetDialect(dialect).Rows(dbApiObject).ToSQL()
		_, err := tx.Exec(query)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Update(ctx context.Context, obj T) error {
	err := withTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		newApiObject, err := toApiObject[T](r, obj)
		if err != nil {
			return err
		}

		query, _, _ := goqu.Update(tableApiObject).SetDialect(dialect).Set(newApiObject).Where(goqu.And(goqu.C(fieldName).Eq(newApiObject.Name), goqu.C(fieldRepoKey).Eq(r.uniqueKey))).ToSQL()
		_, err = tx.Exec(query)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Delete(ctx context.Context, name string) error {
	err := withTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Delete(tableApiObject).SetDialect(dialect).Where(goqu.And(goqu.C(fieldName).Eq(name), goqu.C(fieldRepoKey).Eq(r.uniqueKey))).ToSQL()
		_, err := tx.Exec(query)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Truncate(ctx context.Context) error {
	err := withTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Delete(tableApiObject).SetDialect(dialect).Where(goqu.C(fieldRepoKey).Eq(r.uniqueKey)).ToSQL()
		_, err := tx.Exec(query)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Info() *namespaces.GroupVersionKind {
	return r.gvk
}

type mysqlClusterRepositoryFactory[T domain_helper.Metadatable] struct {
	objFactory domain_helper.MetadatableFactory[T]
}

func (f *mysqlClusterRepositoryFactory[T]) Create(gvk *namespaces.GroupVersionKind, uniqueKeySuffix string) ClusterRepository[T] {
	return NewMysqlClusterRepository(gvk, uniqueKeySuffix, f.objFactory)
}

func newMysqlClusterRepositoryFactory[T domain_helper.Metadatable](objFactory domain_helper.MetadatableFactory[T]) *mysqlClusterRepositoryFactory[T] {
	return &mysqlClusterRepositoryFactory[T]{objFactory: objFactory}
}
