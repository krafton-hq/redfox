package database_helper

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/xo/dburl"
	"go.uber.org/zap"
)

func NewDatabase(dbUrlStr string, username string, password string) (*sqlx.DB, error) {
	urlStr, err := dburl.Parse(dbUrlStr)
	if err != nil {
		return nil, errors.WrapInvalidArguments(err, fmt.Sprintf("Parse 'dbUrl' Failed value: '%s'", dbUrlStr))
	}
	urlStr.User = url.UserPassword(username, password)

	mysqlDsn, err := dburl.GenMysql(urlStr)
	if err != nil {
		return nil, errors.WrapInvalidArguments(err, fmt.Sprintf("'dbUrl' value: '%s'", dbUrlStr))
	}
	zap.S().Infow("Generated Database DSN", "dsn", mysqlDsn, "rawUrl", dbUrlStr)

	return sqlx.Open("mysql", mysqlDsn)
}
