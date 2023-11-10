package models

type Task struct {
	ID          string `json:"id"` // Corrected the struct tag here
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
