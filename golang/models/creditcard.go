package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/durango/go-credit-card"
)

// CreditCard contains general CreditCard
type CreditCard struct {
	ID           int       `schema:"id"`
	CreatedAt    time.Time `schema:"created"`
	UpdatedAt    time.Time `schema:"updated"`
	NameOnCard   string    `schema:"name_on_card" valid:"required"`              //Jane Doe
	Number       string    `schema:"number" gorm:"default:'0'" valid:"required"` //123456789012
	ExpMonth     string    `schema:"exp_month" valid:"required"`                 //02
	ExpYear      string    `schema:"exp_year" valid:"required"`                  //20
	SecurityCode string    `schema:"security_code" valid:"required"`             //1234
	CardType     string    `schema:"card_type" valid:"required"`                 //VISA
	Friendly     string    `schema:"friendly"`                                   //Jane Doe (123456789012) 02/20
}

// SingleLine returns a formatted single line text representing the Model
func (m *CreditCard) SingleLine() string {
	return fmt.Sprintf("%s (%s) %s/%s",
		m.NameOnCard,
		m.Number,
		m.ExpMonth,
		m.ExpYear,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *CreditCard) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *CreditCard) HTMLView() string {
	return "<div id=\"CreditCardHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *CreditCard) HTMLForm() string {
	return "<div id=\"CreditCardHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *CreditCard) IsValid() error {
	if len(m.Number) < 15 || len(m.Number) > 16 {
		return errors.New("Valid Card number is required")
	}

	//If user entered just 1 digit, lets assume they shortened and add the prepending 0
	if len(m.ExpMonth) == 1 {
		m.ExpMonth = fmt.Sprintf("0%s", m.ExpMonth)
	}

	if len(m.ExpMonth) != 2 {
		return errors.New("Invalid expiration month")
	}

	//TODO: Get current year for min, add 10 years for max

	//If user entered just 2 digits, lets assume they shortened and add the prepending 20
	if len(m.ExpYear) == 2 {
		m.ExpYear = fmt.Sprintf("20%s", m.ExpYear)
	}

	if len(m.ExpYear) != 4 {
		return errors.New("Invalid expiration year")
	}

	if len(m.SecurityCode) < 3 || len(m.SecurityCode) > 4 {
		return errors.New("Invalid security code")
	}

	// Initialize a new card:
	card := creditcard.Card{Number: m.Number, Cvv: m.SecurityCode, Month: m.ExpMonth, Year: m.ExpYear}

	// Retrieve the card's method (which credit card company this card belongs to)
	err := card.Method() // card.Company({Short: "visa", Long: "Visa"})
	if err != nil {
		return errors.New("Invalid Card number was entered")
	} else {
		m.CardType = card.Company.Short
	}

	// Display last four digits
	//lastFour, err := card.LastFour() // 4242

	// Validate the card's number (without capturing)
	//err = card.Validate() // will return an error due to not allowing test cards

	err = card.Validate(true) // this will work though

	return err
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *CreditCard) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
