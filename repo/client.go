package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type ClientRepo struct {
	dbProvider *app.DbProvider
}

func NewClientRepo(dbProvider *app.DbProvider) *ClientRepo {
	repo := new(ClientRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo ClientRepo) GetAllCount() (int, error) {
	query, err := repo.dbProvider.GetSql("client_GetAllCount.sql")
	if err != nil {
		return -1, nil
	}
	var count int
	err = repo.dbProvider.Db.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repo ClientRepo) GetSearchCount(search string) (int, error) {
	query, err := repo.dbProvider.GetSql("client_GetSearchCount.sql")
	if err != nil {
		return -1, err
	}
	var count int
	err = repo.dbProvider.Db.QueryRow(query, search).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repo ClientRepo) GetAll() ([]*model.Client, error) {
	query, err := repo.dbProvider.GetSql("client_GetAll.sql")
	if err != nil {
		return nil, err
	}
	var clients []*model.Client
	rows, err := repo.dbProvider.Db.Query(query)
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

func (repo ClientRepo) GetRange(from int, to int) ([]*model.Client, error) {
	query, err := repo.dbProvider.GetSql("client_GetRange.sql")
	if err != nil {
		return nil, err
	}
	var clients []*model.Client
	rows, err := repo.dbProvider.Db.Query(query, from + 1, to + 1)
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
	query, err := repo.dbProvider.GetSql("client_Search.sql")
	if err != nil {
		return nil, err
	}
	var clients []*model.Client
	rows, err := repo.dbProvider.Db.Query(query, search, offset + 1, offset + limit)
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

