package model

import "time"

//ContactInfo - holds contact information
type Contact struct {
	ID        int `sql:",pk,unique" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Category  string `schema:"category" json:"category"`
	FirstName string `schema:"fname" json:"fname"`
	Email     string `schema:"email" json:"email"`
	Message   string `schema:"message" json:"message"`
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
