package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"user_id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`
	Email            string    `gorm:"uniqueIndex;not null" json:"email"`
	Password         string    `gorm:"not null" json:"password"`
	Role             string    `gorm:"type:varchar(255);not null" json:"role"`
	Provider         string    `gorm:"not null" json:"provider"`
	Photo            string    `gorm:"not null" json:"photo"`
	VerificationCode string
	Verified         bool `gorm:"not null" json:"verified"`
}

type SignUp struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
