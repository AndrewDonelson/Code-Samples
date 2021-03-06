package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Profile contains general Profile
type Profile struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Number    int       `schema:"number" gorm:"default:'0'"`     //123456
	String    string    `schema:"string" valid:"required"`       //Test String
	Toggle    bool      `schema:"toggle" gorm:"default:'false'"` //True/False, On/Off, Yes/No
	Float     float64   `schema:"float" gorm:"default:'0.0'"`    //$12.56, 1,123,456.987654321
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Profile) SingleLine() string {
	return fmt.Sprintf("%d %s, %s, %s",
		m.Number,
		m.String,
		strconv.FormatBool(m.Toggle),
		strconv.FormatFloat(m.Float, 'E', -1, 64),
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Profile) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Profile) HTMLView() string {
	return "<div id=\"ProfileHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Profile) HTMLForm() string {
	return "<div id=\"ProfileHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *Profile) IsValid() error {
	if m.String == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Profile) SetFriendly() {
	m.Friendly = m.SingleLine()
	sl := 80
	if len(m.Friendly) < sl {
		sl = len(m.Friendly)
	}

	m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:sl])
}
