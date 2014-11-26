package model

type LookupData struct {
	Href string `json:"href"`
	Activities []*Activity `json:"economicActivities"`
	Branches []*Branch `json:"branches"`
	Cities []*City `json:"cities"`
	Districts []*District `json:"districts"`
	Regions []*Region `json:"regions"`
}

func NewLookupData() *LookupData {
	return new(LookupData)
}
