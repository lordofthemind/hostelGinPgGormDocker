package models

type LoginCredentialModel struct {
	LoginIdentifier string `json:"login_identifier" binding:"required"`
	Password        string `json:"password" binding:"required"`
}
