package factory

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/OPENCBS/server/config"
)

func GetMsSqlConnection(r *http.Request) (conn *sql.DB, err error) {
	conn = nil
	var config config.Configuration
	err = config.Read()
	if err != nil {
		return
	}

	template := "server=%s;user id=%s;password=%s;database=%s"
	connString := fmt.Sprintf(template, config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Name)
	conn, err = sql.Open("mssql", connString)
	return
}

