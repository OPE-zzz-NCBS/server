package api

import (
	"net/http"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/util"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	repo := iface.NewClientRepo()
	offset := util.GetOffset(r)
	limit := util.GetLimit(r)
	query := r.URL.Query().Get("query")
	var items []*model.Client
	var err error
	if query != "" {
		items, err = repo.Search(query, offset, limit)
	} else {
		items, err = repo.GetAll(offset, limit)
	}
	if err != nil {
		fail(w, err)
		return
	}

	count := -1
	if util.GetIncludeCount(r) {
		if query != "" {
			count, err = repo.GetSearchCount(query)
		} else {
			count, err = repo.GetCount()
		}
		if err != nil {
			fail(w, err)
			return
		}
	}

	clients := new(model.Clients)
	clients.Href = util.GetClientsUrl(r)
	clients.Offset = offset
	clients.Limit = limit
	clients.Count = count
	clients.Items = items

	sendJson(w, clients)
}

