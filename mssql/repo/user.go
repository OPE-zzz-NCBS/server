package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model")

type MsSqlUserRepo struct {
}

func NewMsSqlUserRepo() iface.UserRepo {
	return &MsSqlUserRepo{}
}

func (repo MsSqlUserRepo) Find(db *sql.DB) ([]model.User, error) {
	var id int
	var userName string
	var firstName string
	var lastName string
	var users []model.User

	var sql = "select id, user_name, first_name, last_name from dbo.Users"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &userName, &firstName, &lastName)
		if err != nil {
			return nil, err
		}
		user := model.User{id, userName, firstName, lastName}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

