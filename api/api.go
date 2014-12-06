package api

import (
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
	"compress/gzip"
	"github.com/OPENCBS/server/app"
)

type ApiError struct {
	Message string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
	Url string `json:"url"`
}

type APIRequest struct {
	*http.Request
	*app.DbProvider
}

func (r APIRequest) GetRange() (int, int) {
	var offset int
	var limit int

	text := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(text)
	if err != nil {
		offset = 0
	}

	text = r.URL.Query().Get("limit")
	limit, err = strconv.Atoi(text)
	if err != nil {
		limit = 100
	}

	return offset, limit
}

func (r APIRequest) CanAcceptGzip() bool {
	if contentEncoding, ok := r.Header["Accept-Encoding"]; ok {
		return strings.Contains(contentEncoding[0], "gzip")
	}
	return false
}

func sendJson(w http.ResponseWriter, obj interface{}) {
	sendJsonWithStatus(w, obj, http.StatusOK)
}

func sendJsonWithStatus(w http.ResponseWriter, obj interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(obj)
}

func sendCompressedJson(w http.ResponseWriter, obj interface{}) {
	sendCompressedJsonWithStatus(w, obj, http.StatusOK)
}

func sendCompressedJsonWithStatus(w http.ResponseWriter, obj interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(status)
	gz := gzip.NewWriter(w)
	json.NewEncoder(gz).Encode(obj)
	gz.Close()
}

func sendInternalServerError(w http.ResponseWriter, err error) {
	apiError := &ApiError{"Internal server error", err.Error(), ""}
	sendJsonWithStatus(w, apiError, http.StatusInternalServerError)
}
