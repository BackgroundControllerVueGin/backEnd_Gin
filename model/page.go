package model

type Page struct {
	currentPage int `json:"currentPage"`
}

type NextPage struct {
	DataNumber  int `json:"data_number"`
	CurrentPage int `json:"current_page"`
}
