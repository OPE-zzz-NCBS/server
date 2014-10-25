package model

type User struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type Users struct {
	Href string `json:"href"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Items []*User `json:"items"`
}

func NewUser() *User {
	return new(User)
}

