package model

type Client struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewClient() *Client {
	return new(Client)
}

