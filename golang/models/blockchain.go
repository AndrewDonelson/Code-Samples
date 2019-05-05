package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/AndrewDonelson/multichain-client"
)

// Blockchain contains general Blockchain
type Blockchain struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	Client    *multichain.Client
	Chain     string `schema:"chainname" valid:"required"`
	Host      string `schema:"host" valid:"required" gorm:"default:'localhost'"`
	Port      int    `schema:"port"  valid:"required" gorm:"default:'5002'"`

	//Toggle   bool    `schema:"toggle" gorm:"default:'1'"`  //True/False, On/Off, Yes/No
	//Float    float64 `schema:"float" gorm:"default:'0.0'"` //$12.56, 1,123,456.987654321
	Friendly string `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Blockchain) SingleLine() string {
	return fmt.Sprintf("%s: %s:%d",
		m.Chain,
		m.Host,
		m.Port,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Blockchain) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Blockchain) HTMLView() string {
	return "<div id=\"BlockchainHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Blockchain) HTMLForm() string {
	return "<div id=\"BlockchainHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *Blockchain) IsValid() error {
	if m.Chain == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Blockchain) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
