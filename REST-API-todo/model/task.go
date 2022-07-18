package model

type Task struct {
	Title  string `json:"title,omitempty"`
	Id     int    `json:"id,omitempty"`
	Status bool   `json:"status"`
}
