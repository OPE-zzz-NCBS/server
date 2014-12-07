package model

type LookupData struct {
	Href string `json:"href"`
	//Activities []*Activity `json:"economicActivities"`
	Branches []*Branch `json:"branches"`
	Cities []*City `json:"cities"`
	Districts []*District `json:"districts"`
	Regions []*Region `json:"regions"`
	CustomFields []*CustomField `json:"customFields"`
}

func NewLookupData() *LookupData {
	return new(LookupData)
}

