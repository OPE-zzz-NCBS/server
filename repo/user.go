package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/iface"
	"github.com/OPENCBS/server/model"
)

type UserRepo struct {
	sqlProvider iface.SqlProvider
}

func NewUserRepo(sqlProvider iface.SqlProvider) *UserRepo {
	repo := new(UserRepo)
	repo.sqlProvider = sqlProvider
	return repo
}

func (repo UserRepo) getSql(name string) (string, error) {
	return repo.sqlProvider.GetSql(name)
}

func (repo UserRepo) GetAll(db *sql.DB, offset int, limit int) ([]*model.User, error) {
	query, err := repo.getSql("user_GetAll.sql")
	if err != nil {
		return nil, err
	}
	var users []*model.User
	rows, err := db.Query(query, offset + 1, offset + limit)
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

func (repo UserRepo) GetById(db *sql.DB, id int) (*model.User, error) {
	query, err := repo.getSql("user_GetById.sql")
	if err != nil {
		return nil, err
	}
	user := model.NewUser()
	err = db.QueryRow(query, id).Scan(&user.Username, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}

func (repo UserRepo) GetByUsernameAndPassword(db *sql.DB, username string, password string) (*model.User, error) {
	query, err := repo.getSql("user_GetByUsernameAndPassword.sql")
	if err != nil {
		return nil, err
	}
	user := model.NewUser()
	err = db.QueryRow(query, username, password).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Username = username
	return user, nil
}

