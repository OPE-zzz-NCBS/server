package api

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/OPENCBS/server/app"
)

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

