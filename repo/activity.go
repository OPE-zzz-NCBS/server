package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type ActivityRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo ActivityRepo) GetAll() ([]*model.Activity, error) {
	query := repo.GetSql("activity_GetAll.sql")
	rows, err := repo.Db.Query(query)
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

