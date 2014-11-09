package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
)

type MsSqlClientRepo struct {
}

func NewMsSqlClientRepo() iface.ClientRepo {
	return &MsSqlClientRepo{}
}

func (repo MsSqlClientRepo) GetCount(db *sql.DB) (int, error) {
	query, err := Asset("mssql/repo/sql/client_GetCount.sql")
	if err != nil {
		return -1, err
	}

	var count int
	err = db.QueryRow(string(query)).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repo MsSqlClientRepo) FindAll(db *sql.DB, offset int, limit int) ([]*model.Client, error) {
	query, err := Asset("mssql/repo/sql/client_FindAll.sql")
	if err != nil {
		return nil, err
	}

	var clients []*model.Client
	rows, err := db.Query(string(query), offset + 1, offset + limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		client := model.NewClient()
		err := rows.Scan(&client.Id, &client.Name, &client.Type)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

