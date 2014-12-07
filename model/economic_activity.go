package model

type EconomicActivity struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
	ParentId int `json:"parentId"`
}

func NewEconomicActivity() *EconomicActivity {
	return new(EconomicActivity)
}

