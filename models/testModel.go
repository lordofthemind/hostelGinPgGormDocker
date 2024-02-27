package models

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// type BaseModel struct {
// 	gorm.Model
// 	IsActive  bool `gorm:"not null" json:"is_active" default:"true"`
// 	CreatedBy uint `json:"created_by"`
// 	UpdatedBy uint `json:"updated_by"`
// 	DeletedBy uint `json:"deleted_by"`
// }

// type PersonModel struct {
// 	BaseModel
// 	Username    string    `gorm:"not null;unique" json:"username" binding:"required"`
// 	Email       string    `gorm:"not null;unique" json:"email" validate:"required,email"`
// 	Phone       string    `gorm:"not null;unique" json:"phone" binding:"required,gte=10,lte=13"`
// 	Name        string    `gorm:"not null" json:"name" binding:"required"`
// 	JoiningDate time.Time `json:"joining_date"`
// 	Address     string    `gorm:"not null" json:"address"`
// }

// type StudentModel struct {
// 	PersonModel
// 	Aadhar   string `gorm:"not null;unique" json:"aadhar" binding:"required"`
// 	HostelID uint   `json:"hostel_id"` // Foreign key for HostelModel2

// 	// Define one-to-one relationship with BedModel
// 	Bed BedModel `gorm:"foreignKey:StudentID"`
// }

// type BedModel struct {
// 	BaseModel
// 	HostelID   uint `json:"hostel_id"`  // Foreign key for HostelModel2
// 	StudentID  uint `json:"student_id"` // Foreign key for StudentModel2
// 	IsOccupied bool `json:"is_occupied" binding:"required"`
// }

// type HostelModel struct {
// 	BaseModel
// 	Name          string `gorm:"not null" json:"name" binding:"required"`
// 	Address       string `gorm:"not null" json:"address"`
// 	TotalBeds     int    `gorm:"not null" json:"total_beds" binding:"required"`
// 	AvailableBeds int    `gorm:"not null" json:"available_beds" binding:"required"`
// 	AdminID       uint   `json:"admin_id"`  // Foreign key for AdminModel2
// 	WardenID      uint   `json:"warden_id"` // Foreign key for WardenModel2

// 	// Define one-to-many relationship with BedModel
// 	Beds []BedModel `gorm:"foreignKey:HostelID"`
// }

// type CoordinatorModel struct {
// 	PersonModel
// 	PasswordHash string `gorm:"not null" json:"password_hash" binding:"required"`

// 	// Define one-to-one relationship with HostelModel2
// 	Hostel HostelModel `gorm:"foreignKey:CoordinatorID"`

// 	// Additional fields
// 	Role string `gorm:"not null" json:"role" binding:"required"` // Assuming there is a warden role
// }

// type SuperAdminModel struct {
// 	PersonModel
// 	PasswordHash string `gorm:"not null" json:"password_hash" binding:"required"`

// 	// Define one-to-many relationship with HostelModel2
// 	Hostels []HostelModel `gorm:"foreignKey:AdminID"`

// 	// Additional fields
// 	Role string `gorm:"not null" json:"role" binding:"required"` // Assuming there is an admin role
// }
