package models

type Fact struct {
	Id string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Value string `json:"value"`
	ValueUrl string `json:"siteUrl"`
	Provider string `json:"provider"`
}