package factory

import (
	"fmt"
	"net/http"
	"github.com/OPENCBS/server/config"
)

func GetMsSqlConnectionString(r *http.Request) string {
	var config config.Configuration
	config.Read()
	template := "server=%s;user id=%s;password=%s;database=master"
	return fmt.Sprintf(template, config.Database.Host, config.Database.Username, config.Database.Password)
}

