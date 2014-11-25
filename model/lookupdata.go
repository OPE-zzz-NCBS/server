package model

type LookupData struct {
	Href string `json:"href"`
	Activities []*Activity `json:"economicActivities"`
	Branches []*Branch `json:"branches"`
	Cities []*City `json:"cities"`
}

func NewLookupData() *LookupData {
	return new(LookupData)
}
