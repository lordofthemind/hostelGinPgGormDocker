package models

import "gorm.io/gorm"

type StudentModel struct {
	gorm.Model
	Aadhar   string `gorm:"not null;unique" json:"aadhar" binding:"required"`
	Email    string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Phone    string `gorm:"not null;unique" json:"phone" binding:"required,gte=10,lte=13"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Address  string `gorm:"not null" json:"address"`
	IsActive bool   `gorm:"not null" json:"is_active" default:"true"`
}
