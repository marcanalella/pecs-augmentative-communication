package entity

import (
	"github.com/google/uuid"
)

// User represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// swagger:model
type User struct {
	Base

	// the name for this user
	//
	// required: true
	// max length: 128
	Name string `gorm:"type:varchar(128); null" json:"name"`

	// the surname for this user
	//
	// required: true
	// max length: 128
	Surname    string `gorm:"type:varchar(128); null" json:"surname"`
	Email      string `gorm:"type:varchar(128); not null" validate:"required,email,lowercase" json:"email"`
	Password   string `gorm:"type:varchar(128);" json:"password"`
	ResetToken string `gorm:"type:varchar(128); null" json:"resetToken"`
	Role       Role   `gorm:"type:varchar(128);" json:"role"`
}

type Role string

const (
	ADMIN    Role = "ADMIN"
	CUSTOMER      = "CUSTOMER"
)

// GetID returns the user ID
func (u User) GetID() uuid.UUID {
	return u.ID
}

// GetName returns the user name
func (u User) GetName() string {
	return u.Name
}

// GetEmail returns the user email
func (u User) GetEmail() string {
	return u.Email
}

// GetRole returns the user role
func (u User) GetRole() Role {
	return u.Role
}
