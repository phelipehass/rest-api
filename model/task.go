package model

type Task struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Isconcluded bool   `json:"isconcluded,omitempty"`
}
