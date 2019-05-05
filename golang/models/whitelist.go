package models

import (
	"fmt"
	"time"
)

// Whitelist is used for allowing or ignoring certain approved IP addresses
type WhiteList struct {
	ID             int       `schema:"id"`
	CreatedAt      time.Time `schema:"created"`
	UpdatedAt      time.Time `schema:"updated"`
	AccountID      int       `schema:"account_id"`
	BytesWhitelist []byte    `schema:"bytes_whitelist"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *WhiteList) SingleLine() string {
	return fmt.Sprintf("%s", m.BytesWhitelist)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *WhiteList) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *WhiteList) HTMLView() string {
	return "<div id=\"WhiteListHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *WhiteList) HTMLForm() string {
	return "<div id=\"WhiteListHTMLForm\">{Form Content}</div>"
}
