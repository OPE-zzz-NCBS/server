package api

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/OPENCBS/server/app"
)

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

func fail(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func sendJson(w http.ResponseWriter, obj interface{}) {
	sendJsonWithStatus(w, obj, http.StatusOK)
}

func sendJsonWithStatus(w http.ResponseWriter, obj interface{}, status int) {
	json, _ := json.MarshalIndent(obj, "", "  ")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	w.Write(json)
}

func sendInternalServerError(w http.ResponseWriter, err error) {
	apiError := &app.ApiError{"Internal server error", err.Error(), ""}
	sendJsonWithStatus(w, apiError, http.StatusInternalServerError)
}
