package models

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"_id" gorm:"type:uuid;default:gen_random_uuid()"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal("Could not hash password")
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal("Could not hash password")
		return err
	}
	user.Password = string(hash)
	return nil
}
