package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type ClientRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo ClientRepo) GetCount() (int, error) {
	query := repo.GetSql("client_GetCount.sql")
	var count int
	err := repo.Db.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repo ClientRepo) GetSearchCount(search string) (int, error) {
	query := repo.GetSql("client_GetSearchCount.sql")
	var count int
	err := repo.Db.QueryRow(query, search).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repo ClientRepo) GetAll(offset int, limit int) ([]*model.Client, error) {
	query := repo.GetSql("client_GetAll.sql")
	var clients []*model.Client
	rows, err := repo.Db.Query(query, offset + 1, offset + limit)
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

func (repo ClientRepo) Search(search string, offset int, limit int) ([]*model.Client, error) {
	query := repo.GetSql("client_Search.sql")
	var clients []*model.Client
	rows, err := repo.Db.Query(query, search, offset + 1, offset + limit)
	if err != nil {
		return nil, err
	}

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

