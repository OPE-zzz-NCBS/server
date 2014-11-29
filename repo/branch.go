package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type BranchRepo struct {
	dbProvider *app.DbProvider
}

func NewBranchRepo(dbProvider *app.DbProvider) *BranchRepo {
	repo := new(BranchRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo BranchRepo) GetAll() ([]*model.Branch, error) {
	query, err := repo.dbProvider.GetSql("branch_GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repo.dbProvider.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []*model.Branch
	for rows.Next() {
		branch := model.NewBranch()
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Code,
			&branch.Description,
			&branch.Address,
		)
		if err != nil {
			return nil, err
		}
		branches = append(branches, branch)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return branches, nil
}

