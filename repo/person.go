package repo

import (
	"database/sql"
	"strings"
	"github.com/OPENCBS/server/model"
)

type PersonRepo struct {
	GetSql func(name string) string
	Db *sql.DB
}

func (repo PersonRepo) GetById(id int) (*model.Person, error) {
	query := repo.GetSql("person_GetById.sql")
	rows, err := repo.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []*model.Field
	for rows.Next() {
		field := model.NewField()
		var extra string
		err := rows.Scan(&field.Name, &field.DataType, &field.Caption, &field.Value, &extra)
		if err != nil {
			return nil, err
		}
		if field.DataType == "LIST" {
			field.Extra = strings.Split(extra, ":")
		}
		fields = append(fields, field)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	person := model.NewPerson()
	person.Id = id
	person.Fields = fields
	return person, nil
}

