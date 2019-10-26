package models

type Answer struct {
	Songs []Song `json:"objects"`
	Count int    `json:"objects_count"`
}
