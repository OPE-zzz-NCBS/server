package model

type LookupData struct {
	Href string `json:"href"`
	Activities []*Activity `json:"economicActivities"`
	Branches []*Branch `json:"branches"`
}

func NewLookupData() *LookupData {
	return new(LookupData)
}
