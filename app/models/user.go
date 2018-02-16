package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	Name           string     `gorm:"size:255" json:"name"`
	Email          string     `gorm:"type:varchar(100);unique_index" json:"email"`
	HashedPassword []byte     `json:"-"`
	Active         bool       `json:"active"`
	FileName       string     `gorm:"size:255" json:"-"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"-"`
}

// SetNewPassword set a new hashsed password to user
func (user *User) SetNewPassword(passwordString string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	user.HashedPassword = bcryptPassword
}
