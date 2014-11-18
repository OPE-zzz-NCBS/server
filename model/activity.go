package model

type Activity struct {
	Href string `json:"href"`
	Id int `json:"id"`
	Name string `json:"name"`
	ParentId int `json:"parentId"`
}

func NewActivity() *Activity {
	return new(Activity)
}

