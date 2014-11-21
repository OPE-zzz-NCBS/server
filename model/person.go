package model

type Person struct {
	Href string `json:"href"`
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	FatherName string `json:"fatherName"`
	Sex string `json:"sex"`
	BirthDate string `json:"birthDate"`
	BirthPlace string `json:"birthPlace"`
	IdentificationData string `json:"identificationData"`
	Nationality string `json:"nationality"`
	ActivityId int `json:"activityId"`
	BranchId int `json:"branchId"`
}

func NewPerson() *Person {
	return new(Person)
}

