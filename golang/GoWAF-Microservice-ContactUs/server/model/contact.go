package model

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.SetTagName("validate")
}


//ContactInfo - holds contact information
type Contact struct {
	ID        int `sql:",pk,unique" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Category  string `schema:"category" json:"category" validate:"required"`
	FirstName string `schema:"fname" json:"fname" validate:"required"`
	Email     string `schema:"email" json:"email" validate:"required,email"`
	Message   string `schema:"message" json:"message" validate:"required"`
}

func NewContact(fName, email, message string) *Contact {
	article := Contact{
		FirstName: fName,
		Email:     email,
		Message:   message,
	}
	return &article
}

//BeforeInsert - set CreatedAt and UpdatedAt
func (m *Contact) BeforeInsert() error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

//BeforeUpdate - update UpdatedAt
func (m *Contact) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
