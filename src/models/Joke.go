package models

type Joke struct {
	Id      string `json:"id"`
	IconUrl string `json:"icon_url"`
	Url     string `json:"url"`
	Value   string `json:"value"`
}
