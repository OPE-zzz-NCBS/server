package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type CityRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo CityRepo) GetAll() ([]*model.City, error) {
	query := repo.GetSql("city_GetAll.sql")
	var cities []*model.City
	rows, err := repo.Db.Query(query)
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

