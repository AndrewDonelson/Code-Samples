package models

import (
	"errors"
	"fmt"
	"time"
)

// CasbinRule is used for defining access to the available controllers via the casbin library
type CasbinRule struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	AccountID string    `schema:"account_id"`
	Type      string    `schema:"type"`
	Role      string    `schema:"role"`
	Domain    string    `schema:"domain"`
	Object    string    `schema:"object"`
	Action    string    `schema:"action"`
	Friendly  string    `schema:"friendly"`
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *CasbinRule) SetFriendly() {
	m.Friendly = m.SingleLine()
	sl := 80
	if len(m.Friendly) < sl {
		sl = len(m.Friendly)
	}

	m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:sl])
}

// SingleLine returns a formatted single line text representing the Model
func (m *CasbinRule) SingleLine() string {
	return fmt.Sprintf("%s: %s, %s",
		m.Role,
		m.Domain,
		m.Action,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *CasbinRule) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *CasbinRule) HTMLView() string {
	return "<div id=\"CasbinHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *CasbinRule) HTMLForm() string {
	return "<div id=\"CasbinHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *CasbinRule) IsValid() error {
	if m.Role != UserAccessGuest && m.Role != UserAccessMember && m.Role != UserAccessEmployee && len(m.Role) < 1 {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// GetRule returns a casbin rule (either Policy or Group) as a string
func (m *CasbinRule) GetRule() []string {
	switch m.Type {
	case "p":
		return []string{m.Role, m.Domain, m.Object, m.Action}
	case "g":
		return []string{m.AccountID, m.Role, m.Domain}
	default:
		return []string{}
	}
}
