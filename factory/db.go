package factory

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/OPENCBS/server/config"
)

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

