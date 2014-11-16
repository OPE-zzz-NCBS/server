package api

import (
	"net/http"
	"strconv"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/util"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	repo := iface.NewUserRepo()
	offset := util.GetOffset(r)
	limit := util.GetLimit(r)
	items, err := repo.GetAll(offset, limit)
	if err != nil {
		fail(w, err)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get(":id")
	id, _ := strconv.Atoi(idString)
	repo := iface.NewUserRepo()
	user, err := repo.GetById(id)
	if err != nil {
		fail(w, err)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	user.Href = util.GetUserUrl(r, user)

	sendJson(w, user)
}

