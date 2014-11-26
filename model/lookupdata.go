package model

type LookupData struct {
	Href string `json:"href"`
	Activities []*Activity `json:"economicActivities"`
	Branches []*Branch `json:"branches"`
	Cities []*City `json:"cities"`
	Districts []*District `json:"districts"`
}

func NewLookupData() *LookupData {
	return new(LookupData)
}
