package models

import (
	"errors"
	"fmt"
	"time"
)

// Category contains general Category
type Category struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Body      string    `schema:"body" valid:"required"` //Test String
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Category) SingleLine() string {
	return fmt.Sprintf("%s",
		m.Body,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Category) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Category) HTMLView() string {
	return "<div id=\"CategoryHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Category) HTMLForm() string {
	return "<div id=\"CategoryHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *Category) IsValid() error {
	if m.Body == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Category) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
