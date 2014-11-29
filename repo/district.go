package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type DistrictRepo struct {
	dbProvider *app.DbProvider
}

func NewDistrictRepo(dbProvider *app.DbProvider) *DistrictRepo {
	repo := new(DistrictRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo DistrictRepo) GetAll() ([]*model.District, error) {
	query, err := repo.dbProvider.GetSql("district_GetAll.sql")
	if err != nil {
		return nil, err
	}
	var districts []*model.District
	rows, err := repo.dbProvider.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		district := model.NewDistrict()
		err := rows.Scan(
			&district.Id,
			&district.Name,
			&district.RegionId,
		)
		if err != nil {
			return nil, err
		}
		districts = append(districts, district)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return districts, nil
}

