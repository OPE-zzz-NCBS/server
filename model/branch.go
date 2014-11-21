package model

type Branch struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Description string `json:"description"`
	Address string `json:"address"`
}

func NewBranch() *Branch {
	return new(Branch)
}

