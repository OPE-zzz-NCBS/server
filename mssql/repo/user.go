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

func (repo MsSqlUserRepo) FindAll(db *sql.DB) ([]model.User, error) {
	var id int
	var userName string
	var firstName string
	var lastName string
	var users []model.User

	var query = "select id, user_name, first_name, last_name from dbo.Users"
	rows, err := db.Query(query)
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

func (repo MsSqlUserRepo) FindById(db *sql.DB, id int) (*model.User, error) {
	var userName string
	var firstName string
	var lastName string

	var query = "select user_name, first_name, last_name from dbo.Users where id = ?"
	err := db.QueryRow(query, id).Scan(&userName, &firstName, &lastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &model.User{id, userName, firstName, lastName}, nil
}

func (repo MsSqlUserRepo) FindByUsernameAndPassword(db *sql.DB, username string, password string) (*model.User, error) {
	var id int
	var firstName string
	var lastName string

	var query = "select id, first_name, last_name from dbo.Users where user_name = ? and user_pass = ?"
	err := db.QueryRow(query, username, password).Scan(&id, &firstName, &lastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &model.User{id, username, firstName, lastName}, nil
}

