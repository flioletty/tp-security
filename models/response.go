package models

type Response struct {
	ID      string
	Status  int
	Message string
	Headers map[string]any
	Body    string
}
