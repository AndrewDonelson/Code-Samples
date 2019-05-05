package models

import (
	"fmt"
	//"github.com/badoux/checkmail"
	"strings"
	"time"

	nsmisc "github.com/NlaakStudiosLLC/GoWAF-Framework/utils/misc"
)

//TODO: Remove this from public repo

//Email contains a breakdown of a email
//TODO: Update to use/integrate "net/mail" and Address
type Email struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Username  string    `schema:"username"` //bob1234
	Domain    string    `schema:"domain"`   //gmail.com
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Email) SingleLine() string {
	return fmt.Sprintf("%s@%s", m.Username, m.Domain)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Email) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Email) HTMLView() string {
	return "<div id=\"EmailHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Email) HTMLForm() string {
	return "<div id=\"EmailHTMLForm\">{Form Content}</div>"
}

// Sanitize strips all leading and trailing whitespace from strings as well as test normalization all model string properties.
func (m *Email) Sanitize() {
	m.Username = strings.ToLower(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Username)))
	m.Domain = strings.ToLower(strings.TrimSpace(nsmisc.StripCtlAndExtFromUTF8(m.Domain)))
	m.Friendly = strings.TrimSpace(m.SingleLine())
}

//IsValid returns error if model is not complete
func (m *Email) IsValid() error {
	//if err := checkmail.ValidateFormat(m.Friendly); err != nil {
	//	fmt.Println("email err", err)
	//	return err
	//}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Email) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}

// Parse takes a email address as a string and parses it into the model
func (m *Email) Parse(e string) {
	atIdx := strings.Index(e, "@")
	dotIdx := strings.LastIndex(e, ".")
	if atIdx != -1 && dotIdx != -1 {
		m.Username = e[0:atIdx]
		m.Domain = e[atIdx+1 : len(e)]
		m.SetFriendly()
	}
}
