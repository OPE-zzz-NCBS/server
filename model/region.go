package model

type Region struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
}

func NewRegion() *Region {
	return new(Region)
}

