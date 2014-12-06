package api

import (
	"net/http"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/repo"
)

func GetCustomFields(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	repo := repo.NewCustomFieldRepo(ctx.DbProvider)
	customFields, err := repo.GetAll()
	if err != nil {
		apiError := &ApiError{"Internal server error.", err.Error(), ""}
		sendJsonWithStatus(w, apiError, http.StatusInternalServerError)
		return
	}

	sendJson(w, customFields)
}

