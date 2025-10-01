package models

type Task_Json struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
