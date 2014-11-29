package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type UserRepo struct {
	dbProvider *app.DbProvider
}

func NewUserRepo(dbProvider *app.DbProvider) *UserRepo {
	repo := new(UserRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo UserRepo) GetAll(offset int, limit int) ([]*model.User, error) {
	query, err := repo.dbProvider.GetSql("user_GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repo.dbProvider.Db.Query(query, offset + 1, offset + limit)
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
	query, err := repo.dbProvider.GetSql("user_GetById.sql")
	if err != nil {
		return nil, err
	}
	user := model.NewUser()
	err = repo.dbProvider.Db.QueryRow(query, id).Scan(&user.Username, &user.FirstName, &user.LastName)
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
	query, err := repo.dbProvider.GetSql("user_GetByUsernameAndPassword.sql")
	if err != nil {
		return nil, err
	}
	user := model.NewUser()
	err = repo.dbProvider.Db.QueryRow(query, username, password).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	user.Username = username
	return user, nil
}

