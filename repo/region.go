package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type RegionRepo struct {
	dbProvider *app.DbProvider
}

func NewRegionRepo(dbProvider *app.DbProvider) *RegionRepo {
	repo := new(RegionRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo RegionRepo) GetAll() ([]*model.Region, error) {
	query, err := repo.dbProvider.GetSql("region_GetAll.sql")
	if err != nil {
		return nil, err
	}
	var regions []*model.Region
	rows, err := repo.dbProvider.Db.Query(query)
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

