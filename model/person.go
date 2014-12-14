package model

type Address struct {
	CityId int `json:"cityId"`
	Address string `json:"address"`
	PostalCode string `json:"postalCode"`
}

type Person struct {
	Id int `json:"id"`
	UUID string  `json:"uuid"`
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
	HomePhone string `json:"homePhone"`
	PersonalPhone string `json:"personalPhone"`
	Email string `json:"email"`
	Address1 *Address `json:"address1"`
	Address2 *Address `json:"address2"`
	CustomInformation []*CustomFieldValue `json:"customInformation"`
}

func NewPerson() *Person {
	person := new(Person)
	person.Address1 = new(Address)
	person.Address2 = new(Address)
	return person
}
