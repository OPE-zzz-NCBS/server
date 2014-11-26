package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type RegionRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo RegionRepo) GetAll() ([]*model.Region, error) {
	query := repo.GetSql("region_GetAll.sql")
	var regions []*model.Region
	rows, err := repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		region := model.NewRegion()
		err := rows.Scan(&region.Id, &region.Name)
		if err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return regions, nil
}

