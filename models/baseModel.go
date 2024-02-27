package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	IsActive  bool `gorm:"not null;default:false" json:"is_active"`
	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
}

type PersonModel struct {
	BaseModel
	Username    string    `gorm:"not null;unique" json:"username" binding:"required"`
	Email       string    `gorm:"not null;unique" json:"email" validate:"required,email"`
	Phone       string    `gorm:"not null;unique" json:"phone" binding:"required,min=10,max=13"`
	Name        string    `gorm:"not null" json:"name" binding:"required"`
	JoiningDate time.Time `json:"joining_date"`
	Address     string    `gorm:"not null" json:"address"`
}
