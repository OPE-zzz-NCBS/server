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

	var config config.Configuration
	err := config.Read()
	if err != nil {
		return nil, err
	}

	template := "server=%s;user id=%s;password=%s;database=%s"
	connString := fmt.Sprintf(template, config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Name)
	db, err = sql.Open("mssql", connString)
	return db, err
}

