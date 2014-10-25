package util

import (
	"net/http"
	"fmt"
	"strconv"
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

func GetOffset(r *http.Request) int {
	text := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return offset
}

func GetLimit(r *http.Request) int {
	text := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(text)
	if err != nil {
		return 25
	}
	return limit
}

