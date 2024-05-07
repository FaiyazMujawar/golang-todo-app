package requests

import "time"

type CreateTodoRequest struct {
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Expiry      *time.Time `json:"expiry,omitempty"`
}
