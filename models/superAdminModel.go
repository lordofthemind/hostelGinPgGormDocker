package models

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

type SuperAdminModel struct {
	gorm.Model
	Username     string `gorm:"not null;unique" json:"username" binding:"required"`
	Email        string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Phone        string `gorm:"not null;unique" json:"phone" binding:"required,gte=10,lte=13"`
	Name         string `gorm:"not null" json:"name" binding:"required"`
	PasswordHash string `gorm:"not null" json:"-"`
	Address      string `gorm:"not null" json:"address"`
	IsActive     bool   `gorm:"not null" json:"is_active" default:"true"`
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

type SuperAdminModelInterface interface {
	CreateSuperAdmin()
	GetSuperAdmin()
	UpdateSuperAdmin()
	DeleteSuperAdmin()
}

// package entity

// type Person struct {
// 	FirstName string `json:"firstname" binding:"required"`
// 	LastName  string `json:"lastname" binding:"required"`
// 	Age       int8   `json:"age" binding:"gte=1,lte=130"`
// 	Email     string `json:"email" validate:"required,email"`
// }

// type Video struct {
// 	// Title       string `json:"title" xml:"title" form:"title" validate:"email" binding:"required"`
// 	Title       string `json:"title" binding:"min=2,max=1000" validate:"is-cool"`
// 	Description string `json:"description" binding:"max=2000"`
// 	URL         string `json:"url" binding:"required,url"`
// 	Author      Person `json: "author" binding:"required"`
// }
// type GoUser struct {
// 	Username string `json:"username" binding:"required"`
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required,min=8"`
// }

// type Movie struct {
// 	ID       string `json:"id"`
// 	Title    string `json:"title"`
// 	Director string `json:"director"`
// 	Price    string `json:"price"`
// }

// type User struct {
// 	gorm.Model
// 	Email    string `json:"email" gorm:"unique"`
// 	Password string `json:"password"`
// }
// import "gorm.io/gorm"

// type Post struct {
// 	gorm.Model
// 	Title string
// 	Body  string
// }

// func (s *SuperAdminModel) CreateSuperAdmin() {
// 	// Create a new super admin
// }

// func (s *SuperAdminModel) GetSuperAdmin() {
// 	// Get a super admin
// }

// func (s *SuperAdminModel) UpdateSuperAdmin() {
// 	// Update a super admin
// }

// func (s *SuperAdminModel) DeleteSuperAdmin() {
// 	// Delete a super admin
// }

// Path: models/superAdminModel.go
// Compare this snippet from controllers/superAdminController.go:
// package controllers
//
// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/lordofthemind/hostelGinPgGormDocker/models"
// )
//
// type SuperAdminController struct {
// 	SuperAdminModel models.SuperAdminModelInterface
// }
//
// func (s *SuperAdminController) CreateSuperAdmin(c *gin.Context) {
// 	// Create a new super admin
// }
//
// func (s *SuperAdminController) GetSuperAdmin(c *gin.Context) {
// 	// Get a super admin
// }
//
// func (s *SuperAdminController) UpdateSuperAdmin(c *gin.Context) {
// 	// Update a super admin
// }
//
// func (s *SuperAdminController) DeleteSuperAdmin(c *gin.Context) {
// 	// Delete a super admin
// }
