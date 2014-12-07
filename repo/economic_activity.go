package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type EconomicActivityRepo struct {
	dbProvider *app.DbProvider
}

func NewEconomicActivityRepo(dbProvider *app.DbProvider) *EconomicActivityRepo {
	repo := new(EconomicActivityRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo EconomicActivityRepo) GetAll() ([]*model.EconomicActivity, error) {
	query, err := repo.dbProvider.GetSql("economic_activity_GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repo.dbProvider.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var economicActivities []*model.EconomicActivity
	for rows.Next() {
		economicActivity := model.NewEconomicActivity()
		err := rows.Scan(&economicActivity.Id, &economicActivity.Name, &economicActivity.ParentId)
		if err != nil {
			return nil, err
		}
		economicActivities = append(economicActivities, economicActivity)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return economicActivities, nil
}

