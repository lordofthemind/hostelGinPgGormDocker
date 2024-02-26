package models

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

type WardenModel struct {
	gorm.Model
	Username     string `gorm:"not null;unique" json:"username" binding:"required"`
	Email        string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Phone        string `gorm:"not null;unique" json:"phone" binding:"required,gte=10,lte=13"`
	Name         string `gorm:"not null" json:"name" binding:"required"`
	PasswordHash string `gorm:"not null" json:"password_hash" binding:"required"`
	Address      string `gorm:"not null" json:"address"`
	IsActive     bool   `gorm:"not null" json:"is_active" default:"true"`
}

// Validate custom validation logic for the model
func (wa *WardenModel) Validate() error {
	if err := wa.validateEmail(); err != nil {
		return err
	}

	// Add more custom validation logic as needed

	return nil
}

func (wa *WardenModel) validateEmail() error {
	// Example: Validate email format using a simple regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(wa.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

type WardenModelInterface interface {
	CreateWarden()
	GetWarden()
	UpdateWarden()
	DeleteWarden()
}
