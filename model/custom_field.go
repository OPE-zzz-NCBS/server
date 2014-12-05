package model

type CustomFieldHeader struct {
	Id int `json:"id"`
}

type CustomField struct {
	Id int `json:"id"`
	Caption string `json:"caption"`
	Type string `json:"type"`
	Owner string `json:"owner"`
	Tab string `json:"tab"`
	Unique bool `json:"unique"`
	Mandatory bool `json:"mandatory"`
	Order int `json:"order"`
	Extra string `json:"extra"`
}

type CustomFieldValue struct {
	Field *CustomFieldHeader `json:"field"`
	Value string `json:"value"`
}

func NewCustomField() *CustomField {
	return new(CustomField)
}

func NewCustomFieldHeader(id int) *CustomFieldHeader {
	fieldHeader := new(CustomFieldHeader)
	fieldHeader.Id = id
	return fieldHeader
}

func NewCustomFieldValue(fieldId int, value string) *CustomFieldValue {
	customFieldValue := new(CustomFieldValue)
	customFieldValue.Field = NewCustomFieldHeader(fieldId)
	customFieldValue.Value = value
	return customFieldValue
}
