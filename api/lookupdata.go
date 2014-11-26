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

	branchRepo := iface.NewBranchRepo()
	branches, err := branchRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	cityRepo := iface.NewCityRepo()
	cities, err := cityRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	districtRepo := iface.NewDistrictRepo()
	districts, err := districtRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	lookupData := model.NewLookupData()
	lookupData.Activities = activities
	lookupData.Branches = branches
	lookupData.Cities = cities
	lookupData.Districts = districts

	sendJson(w, lookupData)
}

