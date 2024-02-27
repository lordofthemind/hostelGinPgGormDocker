package models

type StudentModel struct {
	PersonModel
	Aadhar   string    `gorm:"not null;unique" json:"aadhar" binding:"required"`
	HostelID uint      `json:"hostel_id"` // Foreign key for HostelModel
	Bed      *BedModel `json:"bed" gorm:"foreignKey:StudentID"`
}
