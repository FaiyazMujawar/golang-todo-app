package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Note struct {
	ID        uuid.UUID `json:"_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body"`
	UserID    uuid.UUID
	Media     pq.StringArray `gorm:"type:varchar(255)[]"`
	User      User           `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
