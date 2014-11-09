package api

import (
	"encoding/json"
	"net/http"
	"log"
	"time"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/util"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	var js []byte
	var repo iface.ClientRepo
	var items []*model.Client
	var clients *model.Clients
	var offset int
	var limit int

	db, err := iface.GetDb(r)
	if err != nil {
		goto Error
	}

	repo = iface.NewClientRepo()
	offset = util.GetOffset(r)
	limit = util.GetLimit(r)
	items, err = repo.FindAll(db, offset, limit)
	if err != nil {
		goto Error
	}

	clients = new(model.Clients)
	clients.Href = util.GetClientsUrl(r)
	clients.Offset = offset
	clients.Limit = limit
	clients.Items = items

	js, err = json.MarshalIndent(clients, "", "  ")
	if err != nil {
		goto Error
	}

	time.Sleep(1000 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(js)
	return

Error:
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

