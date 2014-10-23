package model

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Href string `json:"href"`
}

func NewUser() *User {
	return new(User)
}

