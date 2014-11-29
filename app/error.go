package app

type ApiError struct {
	Message string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
	Url string `json:"url"`
}
