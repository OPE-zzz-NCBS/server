package repo

import (
	"database/sql"
	"github.com/OPENCBS/server/model"
)

type PersonRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo PersonRepo) GetById(id int) (*model.Person, error) {
	query := repo.GetSql("person_GetById.sql")
	person := model.NewPerson()
	err := repo.Db.QueryRow(query, id).Scan(
		&person.FirstName,
		&person.LastName,
		&person.FatherName,
		&person.Sex,
		&person.BirthDate,
		&person.BirthPlace,
		&person.IdentificationData,
		&person.Nationality,
		&person.ActivityId,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	person.Id = id
	return person, nil
}

