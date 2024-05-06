package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `json:"_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Expiry      *time.Time `json:"expiry,omitempty"`
	IsCompleted bool       `json:"completed" gorm:"default:false"`
	UserID      uuid.UUID
	User        User `json:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
