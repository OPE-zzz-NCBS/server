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
	CityId int `json:"cityId"`
	Address string `json:"address"`
	PostalCode string `json:"postalCode"`
	HomePhone string `json:"homePhone"`
	PersonalPhone string `json:"personalPhone"`
	Email string `json:"email"`
}

func NewPerson() *Person {
	return new(Person)
}

