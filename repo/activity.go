package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type ActivityRepo struct {
	dbProvider *app.DbProvider
}

func NewActivityRepo(dbProvider *app.DbProvider) *ActivityRepo {
	repo := new(ActivityRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo ActivityRepo) GetAll() ([]*model.Activity, error) {
	query, err := repo.dbProvider.GetSql("activity_GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repo.dbProvider.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*model.Activity
	for rows.Next() {
		activity := model.NewActivity()
		err := rows.Scan(&activity.Id, &activity.Name, &activity.ParentId)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return activities, nil
}

