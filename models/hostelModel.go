package models

type HostelModel struct {
	BaseModel
	Name          string      `gorm:"not null" json:"name" binding:"required"`
	Address       string      `gorm:"not null" json:"address"`
	TotalBeds     int         `gorm:"not null" json:"total_beds" binding:"required"`
	AvailableBeds int         `gorm:"not null" json:"available_beds" binding:"required"`
	AdminID       uint        `json:"admin_id"`       // Foreign key for AdminModel
	CoordinatorID uint        `json:"coordinator_id"` // Foreign key for CoordinatorModel
	Beds          []*BedModel `json:"beds" gorm:"foreignKey:HostelID"`
}
