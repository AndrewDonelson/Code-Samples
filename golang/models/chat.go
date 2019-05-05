package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Chat contains general Chat
type Chat struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Number    int       `schema:"number" gorm:"default:'0'"`  //123456
	String    string    `schema:"string" valid:"required"`    //Test String
	Toggle    bool      `schema:"toggle" gorm:"default:'1'"`  //True/False, On/Off, Yes/No
	Float     float64   `schema:"float" gorm:"default:'0.0'"` //$12.56, 1,123,456.987654321
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Chat) SingleLine() string {
	return fmt.Sprintf("%d - %s (%s, %s)",
		m.Number,
		m.String,
		strconv.FormatBool(m.Toggle),
		strconv.FormatFloat(m.Float, 'E', -1, 64),
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Chat) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Chat) HTMLView() string {
	return "<div id=\"ChatHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Chat) HTMLForm() string {
	return "<div id=\"ChatHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *Chat) IsValid() error {
	if m.String == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Chat) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
