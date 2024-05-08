package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `json:"_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Body      string    `json:"body" validate:"required"`
	UserID    uuid.UUID
	User      User `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
