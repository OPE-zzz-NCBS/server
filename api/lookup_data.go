package api

/*
import (
	"net/http"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/repo"
)

func GetLookupData(ctx *app.AppContext, w http.ResponseWriter, r *http.Request) {
	lookupData := model.NewLookupData()
	var err error

	lookupData.Activities, err = repo.NewActivityRepo(ctx.DbProvider).GetAll()
	if err != nil {
		goto Error
	}

	lookupData.Branches, err = repo.NewBranchRepo(ctx.DbProvider).GetAll()
	if err != nil {
		goto Error
	}

	lookupData.Cities, err = repo.NewCityRepo(ctx.DbProvider).GetAll()
	if err != nil {
		goto Error
	}

	lookupData.Districts, err = repo.NewDistrictRepo(ctx.DbProvider).GetAll()
	if err != nil {
		goto Error
	}

	lookupData.Regions, err = repo.NewRegionRepo(ctx.DbProvider).GetAll()
	if err != nil {
		goto Error
	}

	lookupData.CustomFields, err = repo.NewCustomFieldRepo(ctx.DbProvider).GetAll()
	if err != nil {
		goto Error
	}

	sendJson(w, lookupData)
	return

Error:
	sendInternalServerError(w, err)
}
*/
