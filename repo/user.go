package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type UserRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo UserRepo) GetAll(offset int, limit int) ([]*model.User, error) {
	query := repo.GetSql("user_GetAll.sql")
	rows, err := repo.Db.Query(query, offset + 1, offset + limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
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

func (repo UserRepo) GetById(id int) (*model.User, error) {
	query := repo.GetSql("user_GetById.sql")
	user := model.NewUser()
	err := repo.Db.QueryRow(query, id).Scan(&user.Username, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}

func (repo UserRepo) GetByUsernameAndPassword(username string, password string) (*model.User, error) {
	query := repo.GetSql("user_GetByUsernameAndPassword.sql")
	user := model.NewUser()
	err := repo.Db.QueryRow(query, username, password).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Username = username
	return user, nil
}

