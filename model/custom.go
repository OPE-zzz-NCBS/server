package model

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
	Field  *CustomField `json:"field"`
	Value string `json:"value"`
}

func NewCustomField() *CustomField {
	return new(CustomField)
}

func NewCustomFieldValue() *CustomFieldValue {
	value := new(CustomFieldValue)
	value.Field = NewCustomField()
	return value
}


