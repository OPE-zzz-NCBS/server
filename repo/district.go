package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type DistrictRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo DistrictRepo) GetAll() ([]*model.District, error) {
	query := repo.GetSql("district_GetAll.sql")
	var districts []*model.District
	rows, err := repo.Db.Query(query)
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

