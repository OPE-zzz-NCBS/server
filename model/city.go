package model

type City struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
	DistrictId int `json:"districtId"`
}

func NewCity() *City {
	return new(City)
}

