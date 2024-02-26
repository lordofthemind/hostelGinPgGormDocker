package models

import "gorm.io/gorm"

type HostelModel struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name" binding:"required"`
	Address       string `gorm:"not null" json:"address"`
	TotalBeds     int    `gorm:"not null" json:"total_beds" binding:"required"`
	AvailableBeds int    `gorm:"not null" json:"available_beds" binding:"required"`
	IsActive      bool   `gorm:"not null" json:"is_active" default:"true"`
}
