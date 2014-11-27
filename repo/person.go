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
		&person.BranchId,
		&person.HomePhone,
		&person.PersonalPhone,
		&person.Address1.CityId,
		&person.Address1.Address,
		&person.Address1.PostalCode,
		&person.Address2.CityId,
		&person.Address2.Address,
		&person.Address2.PostalCode,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	customInformation, err := repo.getCustomInformation(id)
	if err != nil {
		return nil, err
	}
	person.Id = id
	person.CustomInformation = customInformation
	return person, nil
}

func (repo PersonRepo) getCustomInformation(id int) ([]*model.CustomFieldValue, error) {
	query := repo.GetSql("person_GetCustomInformation.sql")
	var values []*model.CustomFieldValue
	rows, err := repo.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		value := model.NewCustomFieldValue()
		err  = rows.Scan(
			&value.Field.Id,
			&value.Field.Caption,
			&value.Field.Type,
			&value.Field.Owner,
			&value.Field.Tab,
			&value.Field.Unique,
			&value.Field.Mandatory,
			&value.Field.Order,
			&value.Field.Extra,
			&value.Value,
		)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return values, nil
}

