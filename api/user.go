package api

import (
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"

	"github.com/OPENCBS/server/repo"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/util"
	"github.com/OPENCBS/server/app"
)

func GetUsers(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	repo := repo.NewUserRepo(ctx.DbProvider)
	offset := util.GetOffset(r)
	limit := util.GetLimit(r)
	items, err := repo.GetAll(offset, limit)
	if err != nil {
		//fail(w, err)
		return
	}
	for _, user := range items {
		user.Href = util.GetUserUrl(r, user)
	}
	users := new(model.Users)
	users.Href = util.GetUsersUrl(r)
	users.Offset = offset
	users.Limit = limit
	users.Items = items

	sendJson(w, users)
}

func GetUser(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	repo := repo.NewUserRepo(ctx.DbProvider)
	user, err := repo.GetById(id)
	if err != nil {
		//fail(w, err)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	user.Href = util.GetUserUrl(r, user)

	sendJson(w, user)
}

