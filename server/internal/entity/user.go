package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User base model
type User struct {
	UserID   uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Login    string    `json:"login" db:"login" validate:"required,lte=30"`
	Password string    `json:"password" db:"password"`
}

// SanitizePassword Sanitize password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// HashPassword Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePasswords Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// PrepareCreate Prepare user for register
func (u *User) PrepareCreate() error {
	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

