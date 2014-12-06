package api

import (
	"fmt"
	"net/http"
	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

func GetClients(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	from, to, err := app.GetRange("clients", r)
	if err != nil {
		apiError := &ApiError{"Requested range is not valid.", err.Error(), ""}
		sendJsonWithStatus(w, apiError, http.StatusRequestedRangeNotSatisfiable)
		return
	}

	repo := repo.NewClientRepo(ctx.DbProvider)
	query := r.URL.Query().Get("query")
	var clients []*model.Client
	var contentRange string
	var status int
	count := 0

	if query != "" {
		count, err = repo.GetSearchCount(query)
		if err != nil {
			goto Error
		}
		if from == -1 {
			clients, err = repo.Search(query)
		} else {
			clients, err = repo.SearchRange(query, from, to)
		}
		if err != nil {
			goto Error
		}
	} else {
		count, err = repo.GetAllCount()
		if err != nil {
			goto Error
		}
		if from == -1 {
			clients, err = repo.GetAll()
		} else {
			clients, err = repo.GetRange(from, to)
		}
		if err != nil {
			goto Error
		}
	}

	if from == -1 {
		status = http.StatusOK
		from = 0
	} else {
		status = http.StatusPartialContent
	}

	to = from + len(clients) - 1
	contentRange = fmt.Sprintf("clients %d..%d/%d", from, to, count)
	w.Header().Set("Accept-Range", "clients")
	w.Header().Set("Content-Range", contentRange)
	sendJsonWithStatus(w, clients, status)
	return

Error:
	sendInternalServerError(w, err)
}

