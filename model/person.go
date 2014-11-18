package model

type Field struct {
	Name string `json:"name"`
	DataType string `json:"dataType"`
	Caption string `json:"caption"`
	Value string `json:"value"`
	Extra interface{} `json:"extra"`
}

type Person struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Fields []*Field `json:"fields"`
}

func NewField() *Field {
	return new(Field)
}

func NewPerson() *Person {
	return new(Person)
}

