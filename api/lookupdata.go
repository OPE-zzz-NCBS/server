package api

import (
	"net/http"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/repo"
)

func GetLookupData(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	activityRepo := repo.NewActivityRepo(ctx.DbProvider)
	activities, err := activityRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	branchRepo := repo.NewBranchRepo(ctx.DbProvider)
	branches, err := branchRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	cityRepo := repo.NewCityRepo(ctx.DbProvider)
	cities, err := cityRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	districtRepo := repo.NewDistrictRepo(ctx.DbProvider)
	districts, err := districtRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	regionRepo := repo.NewRegionRepo(ctx.DbProvider)
	regions, err := regionRepo.GetAll()
	if err != nil {
		fail(w, err)
		return
	}

	lookupData := model.NewLookupData()
	lookupData.Activities = activities
	lookupData.Branches = branches
	lookupData.Cities = cities
	lookupData.Districts = districts
	lookupData.Regions = regions

	sendJson(w, lookupData)
}

