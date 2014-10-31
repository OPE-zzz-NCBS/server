package model

const (
	Person = "PERSON"
	Company = "COMPANY"
	Group = "GROUP"
	VillageBank = "VILLAGE_BANK"
)

type Client struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Clients struct {
	Href string `json:"href"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Items []*Client `json:"items"`
}

func NewClient() *Client {
	return new(Client)
}

