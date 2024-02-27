package models

type BedModel struct {
	BaseModel
	HostelID   uint `json:"hostel_id"`                     // Foreign key for HostelModel
	StudentID  uint `json:"student_id" gorm:"uniqueIndex"` // Foreign key for StudentModel
	IsOccupied bool `json:"is_occupied" binding:"required"`
}
