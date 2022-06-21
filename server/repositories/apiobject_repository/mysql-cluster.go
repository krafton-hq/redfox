package apiobject_repository

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/pkg/transactional"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/json"
)

const tableApiObject = "api_object"
const fieldName = "name"
const fieldRepoKey = "repo_key"

type ApiObject struct {
	Uid        int    `db:"uid" goqu:"skipinsert,skipupdate"`
	RepoKey    string `db:"repo_key" goqu:"skipupdate"`
	Name       string `db:"name" goqu:"skipupdate"`
	ObjectJson string `db:"object_json"`
	Labels     Labels `db:"labels"`
}

type Labels []string

func (l *Labels) Scan(v any) error {
	switch vv := v.(type) {
	case []byte:
		return json.Unmarshal(vv, l)
	case string:
		return json.Unmarshal([]byte(vv), l)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (l Labels) Value() (driver.Value, error) {
	buf, err := json.Marshal(l)
	return buf, err
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

	tr transactional.Transactional
}

func NewMysqlClusterRepository[T domain_helper.Metadatable](gvk *namespaces.GroupVersionKind, uniqueKeySuffix string, factory domain_helper.MetadatableFactory[T], tr transactional.Transactional) *MysqlClusterRepository[T] {
	return &MysqlClusterRepository[T]{
		gvk:       proto.Clone(gvk).(*namespaces.GroupVersionKind),
		uniqueKey: fmt.Sprintf("%s/%s/%s", gvk.Group, gvk.Kind, uniqueKeySuffix),
		factory:   factory,
		tr:        tr,
	}
}

func (r *MysqlClusterRepository[T]) Get(ctx context.Context, name string) (T, error) {
	dbApiObject := &ApiObject{}

	err := r.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Select("*").SetDialect(dialect).From(tableApiObject).Where(goqu.And(goqu.C(fieldName).Eq(name), goqu.C(fieldRepoKey).Eq(r.uniqueKey))).ToSQL()
		row := tx.QueryRowxContext(ctx, query)

		return row.StructScan(dbApiObject)
	})
	if err != nil {
		var t T
		return t, err
	}

	return fromDbApiObject[T](r, dbApiObject)
}

func (r *MysqlClusterRepository[T]) List(ctx context.Context, labelSelectors map[string]string) ([]T, error) {
	lo.Keys[string, string](labelSelectors)

	var dbApiObjects []*ApiObject
	err := r.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Select("*").SetDialect(dialect).From(tableApiObject).
			Where(goqu.And(goqu.C(fieldRepoKey).Eq(r.uniqueKey), goqu.L("? member of ", labelSelectors))).ToSQL()
		zap.S().Info(query)
		return tx.SelectContext(ctx, &dbApiObjects, query)
	})
	if err != nil {
		return nil, errors.WrapInternalError(err, "Sql Error")
	}

	ts := lo.Map[*ApiObject, T](dbApiObjects, func(apiObject *ApiObject, _ int) T {
		var t T
		t, err = fromDbApiObject[T](r, apiObject)
		return t
	})
	if err != nil {
		return nil, err
	}

	var ret []T
	for _, obj := range ts {
		if containsLabels(obj.GetMetadata(), labelSelectors) {
			ret = append(ret, obj)
		}
	}

	return ret, nil
}

func (r *MysqlClusterRepository[T]) Create(ctx context.Context, obj T) error {
	dbApiObject, err := toApiObject[T](r, obj)
	if err != nil {
		return err
	}
	err = r.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Insert(tableApiObject).SetDialect(dialect).Rows(dbApiObject).ToSQL()
		_, err := tx.ExecContext(ctx, query)
		if err != nil {
			zap.S().Infow("Create ApiObject Failed", "query", query)
			return err
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Update(ctx context.Context, obj T) error {
	err := r.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		newApiObject, err := toApiObject[T](r, obj)
		if err != nil {
			return err
		}

		query, _, _ := goqu.Update(tableApiObject).SetDialect(dialect).Set(newApiObject).Where(goqu.And(goqu.C(fieldName).Eq(newApiObject.Name), goqu.C(fieldRepoKey).Eq(r.uniqueKey))).ToSQL()
		result, err := tx.ExecContext(ctx, query)
		if err != nil {
			return err
		}
		if num, _ := result.RowsAffected(); num == 0 {
			return fmt.Errorf("no Rows are Updated %v", obj)
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Delete(ctx context.Context, name string) error {
	err := r.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Delete(tableApiObject).SetDialect(dialect).Where(goqu.And(goqu.C(fieldName).Eq(name), goqu.C(fieldRepoKey).Eq(r.uniqueKey))).ToSQL()
		result, err := tx.ExecContext(ctx, query)
		if err != nil {
			return err
		}
		if num, _ := result.RowsAffected(); num == 0 {
			return errors.NewNotFound(r.uniqueKey + "#" + name)
		}
		return nil
	})
	return err
}

func (r *MysqlClusterRepository[T]) Truncate(ctx context.Context) error {
	err := r.tr.WithTransaction(ctx, func(ctx context.Context, tx *sqlx.Tx, dialect goqu.SQLDialect) error {
		query, _, _ := goqu.Delete(tableApiObject).SetDialect(dialect).Where(goqu.C(fieldRepoKey).Eq(r.uniqueKey)).ToSQL()
		_, err := tx.ExecContext(ctx, query)
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
	tr         transactional.Transactional
}

func (f *mysqlClusterRepositoryFactory[T]) Create(gvk *namespaces.GroupVersionKind, uniqueKeySuffix string) ClusterRepository[T] {
	return NewMysqlClusterRepository(gvk, uniqueKeySuffix, f.objFactory, f.tr)
}

func NewMysqlClusterRepositoryFactory[T domain_helper.Metadatable](objFactory domain_helper.MetadatableFactory[T], tr transactional.Transactional) *mysqlClusterRepositoryFactory[T] {
	return &mysqlClusterRepositoryFactory[T]{objFactory: objFactory, tr: tr}
}
