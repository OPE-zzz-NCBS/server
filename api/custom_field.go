package api

import (
	"net/http"
	"github.com/OPENCBS/server/repo"
)

func GetCustomFields(w http.ResponseWriter, r *APIRequest) {
	repo := repo.NewCustomFieldRepo(r.DbProvider)
	customFields, err := repo.GetAll()
	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	sendJson(w, customFields)
}
