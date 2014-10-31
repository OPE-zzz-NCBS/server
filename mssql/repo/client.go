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

func (repo MsSqlClientRepo) FindAll(db *sql.DB, offset int, limit int) ([]*model.Client, error) {
	var clients []*model.Client

	var query =
		"select id, name, client_type " +
		"from (" +
		"	select id, first_name + ' ' + last_name name, 'PERSON' client_type, " +
		"		row_number() over (order by id asc) num " +
		"	from dbo.Persons " +
		") t " +
		"where t.num between ? and ?"
	rows, err := db.Query(query, offset + 1, offset + limit)
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

