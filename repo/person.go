package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type PersonRepo struct {
	dbProvider *app.DbProvider
}

func NewPersonRepo(dbProvider *app.DbProvider) *PersonRepo {
	repo := new(PersonRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo PersonRepo) GetPeople(offset int, limit int) ([]*model.Person, error) {
	query, err := repo.dbProvider.GetSql("person_GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repo.dbProvider.Db.Query(query, offset + 1, offset + limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var person *model.Person
	var people []*model.Person
	for rows.Next() {
		thisPerson := model.NewPerson()
		var customFieldId int
		var customFieldValue string
		err := rows.Scan(
			&thisPerson.Id,
			&thisPerson.FirstName,
			&thisPerson.LastName,
			&thisPerson.FatherName,
			&thisPerson.Sex,
			&thisPerson.BirthDate,
			&thisPerson.BirthPlace,
			&thisPerson.IdentificationData,
			&thisPerson.Nationality,
			&thisPerson.ActivityId,
			&thisPerson.BranchId,
			&thisPerson.HomePhone,
			&thisPerson.PersonalPhone,
			&thisPerson.Address1.CityId,
			&thisPerson.Address1.Address,
			&thisPerson.Address1.PostalCode,
			&thisPerson.Address2.CityId,
			&thisPerson.Address2.Address,
			&thisPerson.Address2.PostalCode,
			&customFieldId,
			&customFieldValue,
		)
		if err != nil {
			return nil, err
		}
		if person == nil || person.Id != thisPerson.Id {
			if person != nil {
				people = append(people, person)
			}
			person = thisPerson
		}
		if customFieldId > 0 {
			value := model.NewCustomFieldValue(customFieldId, customFieldValue)
			person.CustomInformation = append(person.CustomInformation, value)
		}
	}
	if person != nil {
		people = append(people, person)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (repo PersonRepo) Add(person *model.Person) (*model.Person, error) {
	tx, err := repo.dbProvider.Db.Begin()
	if err != nil {
		return nil, err
	}
	query, err := repo.dbProvider.GetSql("client_AddTiers.sql")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	res, err := stmt.Exec(
		person.Address1.CityId,
		person.Address2.CityId,
		person.BranchId,
		person.HomePhone,
		person.PersonalPhone,
		person.Email,
		person.Address1.Address,
		person.Address1.PostalCode,
		person.Address2.Address,
		person.Address2.PostalCode,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	query, err = repo.dbProvider.GetSql("person_Add.sql")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	stmt, err = tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	person.Id = int(id)
	_, err = stmt.Exec(
		person.Id,
		person.FirstName,
		person.LastName,
		person.FatherName,
		person.Sex,
		person.BirthDate,
		person.BirthPlace,
		person.IdentificationData,
		person.Nationality,
		person.ActivityId,
	)
	if err != nil {
		person.Id = 0
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return person, nil
}

/*
func (repo PersonRepo) GetById(id int) (*model.Person, error) {
	query, err := repo.dbProvider.GetSql("person_GetById.sql")
	if err != nil {
		return nil, err
	}
	person := model.NewPerson()
	err = repo.dbProvider.Db.QueryRow(query, id).Scan(
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
	query, err := repo.dbProvider.GetSql("person_GetCustomInformation.sql")
	if err != nil {
		return nil, err
	}
	var values []*model.CustomFieldValue
	rows, err := repo.dbProvider.Db.Query(query, id)
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
*/