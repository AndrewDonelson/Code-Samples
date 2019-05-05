package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	nsmisc "github.com/NlaakStudiosLLC/GoWAF-Framework/utils/misc"
)

//PersonName hold the complete name of a person
type PersonName struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Prefix    string    `schema:"prefix"`  //ie. Mr
	First     string    `schema:"first"`   //William
	Middle    string    `schema:"middle"`  //Blaine
	Last      string    `schema:"last"`    //Doe
	Suffix    string    `schema:"suffix"`  //Sr
	GoesBy    string    `schema:"goes_by"` //Bob
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *PersonName) SingleLine() string {
	return fmt.Sprintf("%s %s %s %s %s",
		m.Prefix,
		m.First,
		m.Middle,
		m.Last,
		m.Suffix,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *PersonName) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *PersonName) HTMLView() string {
	return "<div id=\"PersonNameHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *PersonName) HTMLForm() string {
	return "<div id=\"PersonNameHTMLForm\">{Form Content}</div>"
}

// Sanitize strips all leading and trailing whitespace from strings as well as test normalization all model string properties.
func (m *PersonName) Sanitize() {
	m.Prefix = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Prefix)))
	m.First = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.First)))
	m.Middle = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Middle)))
	m.Last = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Last)))
	m.Suffix = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Suffix)))
	m.GoesBy = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.GoesBy)))
	m.Friendly = strings.TrimSpace(m.SingleLine())
}

//IsValid returns error if model is not complete
func (m *PersonName) IsValid() error {
	if m.First == "" || m.Last == "" {
		return errors.New("first or last name can't be empty")
	}

	return nil
}

// Parse takes a given strings and tried to break it up into name parts and assigns to object
func (m *PersonName) Parse(s string) {
	parts := strings.Fields(s)
	if len(parts) == 2 {
		m.First = parts[0]
		m.Last = parts[1]
	} else if len(parts) == 3 {
		m.First = parts[0]
		m.Middle = parts[1]
		m.Last = parts[2]
	}
	m.SetFriendly()
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *PersonName) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
