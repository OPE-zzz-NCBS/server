package repo

import (
	"github.com/OPENCBS/server/model"
	"github.com/OPENCBS/server/app"
)

type CustomFieldRepo struct {
	dbProvider *app.DbProvider
}

func NewCustomFieldRepo(dbProvider *app.DbProvider) *CustomFieldRepo{
	repo := new(CustomFieldRepo)
	repo.dbProvider = dbProvider
	return repo
}

func (repo CustomFieldRepo) GetAll() ([]*model.CustomField, error) {
	query, err := repo.dbProvider.GetSql("custom_field_GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := repo.dbProvider.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customFields []*model.CustomField
	for rows.Next() {
		customField := model.NewCustomField()
		err := rows.Scan(
			&customField.Id,
			&customField.Caption,
			&customField.Type,
			&customField.Owner,
			&customField.Tab,
			&customField.Unique,
			&customField.Mandatory,
			&customField.Order,
			&customField.Extra,
		)
		if err != nil {
			return nil, err
		}
		customFields = append(customFields, customField)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return customFields, nil
}

