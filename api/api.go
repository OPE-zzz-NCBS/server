package api

import (
	"log"
	"net/http"
	"encoding/json"
)

func fail(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func sendJson(w http.ResponseWriter, obj interface{}) {
	js, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fail(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
}

