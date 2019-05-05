package models

import (
	"errors"
	"fmt"
	"time"
)

// Tag contains general Tag
type Tag struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Body      string    `schema:"body" valid:"required"`
	TagID		int 	`schema:"tag_id"`
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Tag) SingleLine() string {
	return fmt.Sprintf("%s",
		m.Body,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Tag) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Tag) HTMLView() string {
	return "<div id=\"TagHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Tag) HTMLForm() string {
	return "<div id=\"TagHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *Tag) IsValid() error {
	if m.Body == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Tag) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
