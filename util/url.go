package util

import (
	"fmt"
	"net/http"
	"github.com/OPENCBS/server/model"
)

func GetBaseUrl(r *http.Request) string {
	var schema string
	if r.TLS != nil {
		schema = "https"
	} else {
		schema = "http"
	}
	return fmt.Sprintf("%s://%s/api", schema, r.Host)
}

func GetUsersUrl(r *http.Request) string {
	return fmt.Sprintf("%s/users", GetBaseUrl(r))
}

func GetUserUrl(r *http.Request, u *model.User) string {
	return fmt.Sprintf("%s/users/%d", GetBaseUrl(r), u.Id)
}

func GetClientsUrl(r *http.Request) string {
	return fmt.Sprintf("%s/clients", GetBaseUrl(r))
}

