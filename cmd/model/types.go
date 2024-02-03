package model

import (
	"html/template"
)

type State struct {
	Id        int    `json:"id"`
	StateCode string `json:"stateCode"`
	StateName string `json:"stateName"`
	EpaRegion string `json:"epaRegion"`
}

type Data struct {
	Contacts Contacts
	States   []State
}
type Templates struct {
	templates *template.Template
}
type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
}

type Contacts = []Contact
