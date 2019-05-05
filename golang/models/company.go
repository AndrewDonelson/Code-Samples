package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	nsmisc "github.com/NlaakStudiosLLC/GoWAF-Framework/utils/misc"
)

//Company stores information about the company
type Company struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Name      string    `schema:"name"`
	ContactID int       `schema:"contact_id"`
	Person    Person    `schema:"contactperson" gorm:"foreignkey:ContactID" gorm:"auto_preload"`
	AddressID int       `schema:"address_id"`
	Address   Address   `schema:"address" gorm:"foreignkey:AddressID" gorm:"auto_preload"`
	EmailID   int       `schema:"email_id"`
	Email     Email     `schema:"email" gorm:"foreignkey:EmailID" gorm:"auto_preload"`
	PhoneID   int       `schema:"phone_id"`
	Phone     Phone     `schema:"phone" gorm:"foreignkey:PhoneID" gorm:"auto_preload"`
	FaxID     int       `schema:"fax_id"`
	Fax       Phone     `schema:"fax" gorm:"foreignkey:FaxID" gorm:"auto_preload"`
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Company) SingleLine() string {
	if m.ContactID != 0 {
		return fmt.Sprintf("%s, PH: %s - POC: %s", m.Name, m.Phone.Friendly, m.Person.Friendly)
	}
	return fmt.Sprintf("%s, POC: %s", m.Name, m.Person.Friendly)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Company) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Company) HTMLView() string {
	return "<div id=\"CompanyHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Company) HTMLForm() string {
	return "<div id=\"CompanyHTMLForm\">{Form Content}</div>"
}

// Sanitize strips all leading and trailing whitespace from strings as well as test normalization all model string properties.
func (m *Company) Sanitize() {
	m.Name = strings.Title(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Name)))
	m.Friendly = strings.TrimSpace(m.SingleLine())
}

//IsValid returns error if model is not complete
func (m *Company) IsValid() error {
	if err := m.Person.IsValid(); err != nil {
		fmt.Println("1", err)
		return err
	}
	return nil
}

//IsValid returns error if model is not complete
func (m *Company) IsValidName() error {
	if m.Name == "" {
		return errors.New("Must at least a valid company name")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Company) SetFriendly() {
	m.Friendly = m.SingleLine()
	m.Email.SetFriendly()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
