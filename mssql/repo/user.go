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

func (repo MsSqlUserRepo) FindAll(db *sql.DB) ([]*model.User, error) {
	var users []*model.User

	var query = "select id, user_name, first_name, last_name from dbo.Users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := model.NewUser()
		err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo MsSqlUserRepo) FindById(db *sql.DB, id int) (*model.User, error) {
	user := model.NewUser()
	var query = "select user_name, first_name, last_name from dbo.Users where id = ?"
	err := db.QueryRow(query, id).Scan(&user.Username, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}

func (repo MsSqlUserRepo) FindByUsernameAndPassword(db *sql.DB, username string, password string) (*model.User, error) {

	user := model.NewUser()
	var query = "select id, first_name, last_name from dbo.Users where user_name = ? and user_pass = ?"
	err := db.QueryRow(query, username, password).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Username = username
	return user, nil
}

