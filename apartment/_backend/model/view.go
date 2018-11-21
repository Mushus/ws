package model

type LoginView struct {
	login  string
	Errors []error
}

type ApartmentView struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
