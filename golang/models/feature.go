package models

import (
	"fmt"
	"time"
)

type Feature struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	UUID      string    `schema:"uuid" valid:"required"` //UUID V4
	Name      string    `schema:"name" valid:"required"`
	Version   string    `schema:"version" valid:"required"`
	Publisher string    `schema:"publisher" valid:"required"`
	Support   string    `schema:"support" valid:"required"`
	Verbose   bool      `schema:"verbose" gorm:"default:'0'"`  //True/False, On/Off, Yes/No
	NoModel   bool      `schema:"no_model" gorm:"default:'1'"` //True/False, On/Off, Yes/No
	Uses      []string  `schema:"uses"`
	Tags      []string  `schema:"tags"`
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Feature) SingleLine() string {
	return fmt.Sprintf("%s v%s", m.Name, m.Version)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Feature) MultiLine() string {
	return m.SingleLine()
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Feature) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
