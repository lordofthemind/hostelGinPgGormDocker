package models

import "gorm.io/gorm"

type WardenModel struct {
	gorm.Model
	Name         string
	Email        string
	PasswordHash string
	PasswordSalt string
	IsActive     bool
}

type WardenModelInterface interface {
	CreateWarden()
	GetWarden()
	UpdateWarden()
	DeleteWarden()
}
