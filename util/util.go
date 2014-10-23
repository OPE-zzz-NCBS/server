package util

import (
	"net/http"
	"fmt"
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

