package model

type District struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
	RegionId int `json:"regionId"`
}

func NewDistrict() *District {
	return new(District)
}

