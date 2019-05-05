package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	nsmisc "github.com/NlaakStudiosLLC/GoWAF-Framework/utils/misc"
	"github.com/darkoatanasovski/htmltags"
)

// Blog
type Blog struct {
	ID         int       `schema:"id"`
	CreatedAt  time.Time `schema:"created"`
	UpdatedAt  time.Time `schema:"updated"`
	PersonID   int       `schema:"person_id"`
	Person     Person    `gorm:"foreignkey:PersonID"`
	Title      string    `schema:"title"`
	Body       string    `schema:"body"`
	Tag        []Tag     `gorm:"foreignkey:TagID"`
	CategoryID int       `schema:"category_id"`
	Category   Category  `gorm:"foreignkey:CategoryID"`
	Friendly   string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Blog) SingleLine() string {
	sl := 40
	if len(m.Body) < 40 {
		sl = len(m.Body)
	}
	stripped, _ := htmltags.Strip(m.Body[0:sl], []string{}, false)
	return fmt.Sprintf("%s..., (%s)",stripped.ToString(), m.Person.PrimaryName.SingleLine())
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Blog) MultiLine() string {
	return fmt.Sprintf("%s:\n%s", m.Person.PrimaryName.SingleLine(), m.Body)
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Blog) HTMLView() string {
	return "<div id=\"BlogHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Blog) HTMLForm() string {
	return "<div id=\"BlogHTMLForm\">{Form Content}</div>"
}

// Sanitize strips all leading and trailing whitespace from strings as well as test normalization all model string properties.
func (m *Blog) Sanitize() {
	m.Body = strings.ToLower(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Body)))
	m.Friendly = strings.TrimSpace(m.SingleLine())
}

//IsValid returns error if model is not complete
func (m *Blog) IsValid() error {
	if m.Body == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Blog) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
