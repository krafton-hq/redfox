package apiobject_repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/domain_helper"
	"github.com/krafton-hq/red-fox/server/pkg/transactional"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlClusterRepo *MysqlClusterRepository[*documents.NatIp]
var db *sqlx.DB
var tr transactional.Transactional

func init() {
	db2, err := sqlx.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/red_fox")
	if err != nil {
		log.Fatal(err)
	}
	db = db2

	tr, err = transactional.NewSqlTransactional(db, "mysql", nil)
	if err != nil {
		log.Fatal(err)
	}

	mysqlClusterRepo = NewMysqlClusterRepository[*documents.NatIp](domain_helper.NatIpGvk, "default", domain_helper.NewNatIpFactory(), tr)
}

func TestMysqlClusterRepository_Create(t *testing.T) {
	gvk := domain_helper.NatIpGvk

	obj := &documents.NatIp{
		ApiVersion: gvk.Group + "/" + gvk.Version,
		Kind:       gvk.Kind,
		Metadata: &idl_common.ObjectMeta{
			Name:      "my-first-object4",
			Labels:    nil,
			Namespace: "default",
		},
		Spec: &documents.NatIpSpec{
			Type: documents.IpType_Ipv4,
			Cidrs: []string{
				"1.1.1.1/32",
				"2.2.2.2/32",
			},
		},
	}

	ctx := tr.WithDatabaseContext(context.TODO())
	err := mysqlClusterRepo.Create(ctx, obj)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMysqlClusterRepository_Get(t *testing.T) {
	ctx := tr.WithDatabaseContext(context.TODO())
	natIp, err := mysqlClusterRepo.Get(ctx, "my-first-object4")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", natIp)
}

func TestMysqlClusterRepository_Update(t *testing.T) {
	gvk := domain_helper.NatIpGvk
	obj := &documents.NatIp{
		ApiVersion: gvk.Group + "/" + gvk.Version,
		Kind:       gvk.Kind,
		Metadata: &idl_common.ObjectMeta{
			Name: "my-first-object4",
			Labels: map[string]string{
				"foo": "bar",
			},
			Namespace: "default",
		},
		Spec: &documents.NatIpSpec{
			Type: documents.IpType_Ipv4,
			Cidrs: []string{
				"8.8.8.8/32",
				"8.8.4.4/32",
			},
		},
	}

	ctx := tr.WithDatabaseContext(context.TODO())
	err := mysqlClusterRepo.Update(ctx, obj)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMysqlClusterRepository_Delete(t *testing.T) {
	ctx := tr.WithDatabaseContext(context.TODO())
	err := mysqlClusterRepo.Delete(ctx, "my-first-object2")
	if err != nil {
		t.Fatal(err)
	}
}
