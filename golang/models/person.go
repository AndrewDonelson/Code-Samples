package models

import (
	"fmt"
	"strings"
	"time"
)

//Person contains all data pertaining to a individual person
type Person struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	//Dob       time.Time `schema:"dob" json:"date,omitempty"`

	//GenderID is the UID of the Person's Gender as found in the gender table
	GenderID int    `schema:"gender_id"`
	Gender   Gender `gorm:"foreignkey:GenderID" gorm:"auto_preload"`

	//NameID is the UID of the Person's name as found in the person_name table
	NameID      int        `schema:"name_id"`
	PrimaryName PersonName `gorm:"foreignkey:NameID" gorm:"auto_preload"`

	//SpouseNameID is the UID of the Person's name as found in the person_name table
	SpouseNameID int        `schema:"spouse_name_id"`
	SpouseName   PersonName `gorm:"foreignkey:SpouseNameID" gorm:"auto_preload"`

	//EmailID is the UID of the Person's email as found in the email table
	EmailID int   `schema:"email_id"`
	Email   Email `gorm:"foreignkey:EmailID" gorm:"auto_preload"`

	//TypeID is the UID of the Person's Type as found in the person_type table
	TypeID     int        `schema:"type_id"`
	PersonType PersonType `gorm:"foreignkey:TypeID" gorm:"auto_preload"`

	AddressID int     `schema:"address_id"`
	Address   Address `gorm:"foreignkey:AddressID" gorm:"auto_preload"`

	//PhoneID is the UID of the Person's Phone info as found in the phone table
	PhoneID  int    `schema:"phone_id"`
	Phone    Phone  `gorm:"foreignkey:PhoneID" gorm:"auto_preload"`
	Friendly string `schema:"friendly"`
	//CompanyID is the UID of the Company's where person employee
	CompanyID int `schema:"company_id"`
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Person) SetFriendly() {
	m.Friendly = m.SingleLine()
	sl := 80
	if len(m.Friendly) < sl {
		sl = len(m.Friendly)
	}
	m.Phone.SetFriendly()
	m.Gender.SetFriendly()
	m.PersonType.SetFriendly()
	m.PrimaryName.SetFriendly()
	//m.SpouseName.SetFriendly()
	m.Address.SetFriendly()
	m.Email.SetFriendly()
	m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:sl])
}

// SingleLine returns a formatted single line text representing a Person Model
func (m *Person) SingleLine() string {

	return fmt.Sprintf("%s [%s], %s, %s",
		m.PrimaryName.SingleLine(),
		m.PersonType.SingleLine(),
		m.Email.SingleLine(),
		m.Phone.SingleLine(),
	)
}

// MultiLine returns a formatted multi-line text representing a Person Model
func (m *Person) MultiLine() string {
	pn := m.PrimaryName.SingleLine()
	return fmt.Sprintf("%s\n%s\n%s\n%s\n",
		pn,
		m.PersonType.Name,
		m.Email.Friendly,
		m.Phone.Friendly,
	)
}

// HTMLView returns a HTML5 code representing a view of a Person Model
func (m *Person) HTMLView() string {
	return "<div id=\"PersonHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of a Person Model
func (m *Person) HTMLForm() string {
	return "<div id=\"PersonHTMLForm\">{Form Content}</div>"
}

// Sanitize strips all leading and trailing whitespace from strings as well as test normalization all model string properties.
func (m *Person) Sanitize() {
	m.Friendly = strings.TrimSpace(m.SingleLine())
}

//IsValid returns error if model is not complete
func (m *Person) IsValid() error {
	if err := m.PrimaryName.IsValid(); err != nil {
		return err
	}

	if err := m.Email.IsValid(); err != nil {
		return err
	}

	if err := m.Phone.IsValid(); err != nil {
		return err
	}

	if err := m.Gender.IsValid(); err != nil {
		return err
	}

	return nil
}
