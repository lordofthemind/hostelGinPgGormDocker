package models

import (
	"errors"
	"regexp"
)

type CoordinatorModel struct {
	PersonModel
	PasswordHash string       `gorm:"not null" json:"password_hash" binding:"required"`
	Hostel       *HostelModel `json:"hostel" gorm:"foreignKey:CoordinatorID"` // Use CoordinatorID as the foreign key
	Role         string       `gorm:"not null" json:"role" binding:"required"`
}

// Validate custom validation logic for the model
func (co *CoordinatorModel) Validate() error {
	if err := co.validateEmail(); err != nil {
		return err
	}

	// Add more custom validation logic as needed

	return nil
}

func (co *CoordinatorModel) validateEmail() error {
	// Email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(co.Email) {
		return errors.New("invalid email format")
	}
	return nil
}
