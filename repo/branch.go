package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type BranchRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo BranchRepo) GetAll() ([]*model.Branch, error) {
	query := repo.GetSql("branch_GetAll.sql")
	rows, err := repo.Db.Query(query)
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

