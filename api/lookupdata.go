package api

import (
	"net/http"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
)

func GetLookupData(w http.ResponseWriter, r *http.Request) {
	activityRepo := iface.NewActivityRepo()
	activities, err := activityRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	lookupData := model.NewLookupData()
	lookupData.Activities = activities

	sendJson(w, lookupData)
}

