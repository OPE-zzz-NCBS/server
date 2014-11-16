package factory

import (
	"fmt"
	"path"
	"net/http"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/OPENCBS/server/config"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/sql_mssql"
)

type MsSqlSqlProvider struct {
}

var db *sql.DB

func GetMsSqlDb(r *http.Request) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var conf *config.Configuration
	conf, err := config.Get()
	if err != nil {
		return nil, err
	}

	template := "server=%s;user id=%s;password=%s;database=%s;connection timeout=5"
	connString := fmt.Sprintf(template,
		conf.Database.Host,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Name)
	db, err = sql.Open("mssql", connString)
	return db, err
}

func GetDb(r *http.Request) (*sql.DB, error) {
	return GetMsSqlDb(r)
}

func GetSqlProvider(r *http.Request) iface.SqlProvider {
	return new(MsSqlSqlProvider)
}

func (sqlProvider MsSqlSqlProvider) GetSql(name string) (string, error) {
	path := path.Join("sql_mssql", name)
	sql, err := sql_mssql.Asset(path)
	if err != nil {
		return "", err
	}
	return string(sql), nil
}

