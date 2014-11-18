package model

type LookupData struct {
	Href string `json:"href"`
	Activities []*Activity `json:"economicActivities"`
}

func NewLookupData() *LookupData {
	return new(LookupData)
}
