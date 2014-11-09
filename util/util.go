package util

import (
	"net/http"
	"strconv"
)

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

func GetIncludeCount(r *http.Request) bool {
	text := r.URL.Query().Get("includeCount")
	includeCount, err := strconv.ParseBool(text)
	if err != nil {
		return false
	}
	return includeCount
}

