package models

import (
	"errors"
	"regexp"
)

type SuperAdminModel struct {
	PersonModel
	PasswordHash string         `gorm:"not null" json:"password_hash" binding:"required"`
	Hostels      []*HostelModel `json:"hostels" gorm:"foreignKey:AdminID"`
	Role         string         `gorm:"not null" json:"role" binding:"required"`
}

// Validate custom validation logic for the model
func (sa *SuperAdminModel) Validate() error {
	if err := sa.validateEmail(); err != nil {
		return err
	}

	// Add more custom validation logic as needed

	return nil
}

func (sa *SuperAdminModel) validateEmail() error {
	// Example: Validate email format using a simple regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(sa.Email) {
		return errors.New("invalid email format")
	}
	return nil
}
