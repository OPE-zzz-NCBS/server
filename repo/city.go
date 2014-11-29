package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type CityRepo struct {
	dbProvider *app.DbProvider
}

func NewCityRepo(dbProvider *app.DbProvider) *CityRepo {
	repo := new(CityRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo CityRepo) GetAll() ([]*model.City, error) {
	query, err := repo.dbProvider.GetSql("city_GetAll.sql")
	if err != nil {
		return nil, err
	}
	var cities []*model.City
	rows, err := repo.dbProvider.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		city := model.NewCity()
		err := rows.Scan(&city.Id, &city.Name, &city.DistrictId)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return cities, nil
}

