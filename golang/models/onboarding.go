package models

import (
	"errors"
	"fmt"
	"time"
)

// Onboarding contains general Onboarding
type Onboarding struct {
	ID        int       `schema:"id"`
	CreatedAt time.Time `schema:"created"`
	UpdatedAt time.Time `schema:"updated"`
	OwnerID   int       `schema:"owner_id"`                                        // 0 = WebApp or AccountID for User and Companies
	Slide     int       `schema:"slide" valid:"required"`                          // the slide index
	Title     string    `schema:"title" valid:"required"`                          // Slide Title 24 max
	Message   string    `schema:"message" valid:"required"`                        // Slide Message 100 maax
	ImgBGPath string    `schema:"im_bg_path"`                                      // path to slide background image /static/images/slides/bg/...
	Icon      string    `schema:"icon"`                                            // Icon to use
	Finished  string    `schema:"Finished" valid:"required" gorm:"default:'next'"` // Route to redirect to after complete
	Friendly  string    `schema:"friendly"`
}

// SingleLine returns a formatted single line text representing the Model
func (m *Onboarding) SingleLine() string {
	return fmt.Sprintf("%d - %s",
		m.Slide,
		m.Title,
	)
}

// MultiLine returns a formatted multi-line text representing the Model
func (m *Onboarding) MultiLine() string {
	return m.SingleLine()
}

// HTMLView returns a HTML5 code representing a view of the Model
func (m *Onboarding) HTMLView() string {
	return "<div id=\"OnboardingHTMLView\">{View Content}</div>"
}

// HTMLForm returns a HTML5 code representing a form of the Model
func (m *Onboarding) HTMLForm() string {
	return "<div id=\"OnboardingHTMLForm\">{Form Content}</div>"
}

//IsValid returns error if model is not complete
func (m *Onboarding) IsValid() error {
	if m.Title == "" || m.Slide == 0 || m.Message == "" || m.Finished == "" {
		return errors.New("Please fill in all required fields")
	}
	return nil
}

// SetFriendly sets the Friendly (short summary) of model data. This should be called on every Create and Update for every model.
func (m *Onboarding) SetFriendly() {
	m.Friendly = m.SingleLine()
	maxlen := 80
	if len(m.Friendly) <= maxlen {
		m.Friendly = fmt.Sprintf("%s", m.Friendly)
	} else {
		m.Friendly = fmt.Sprintf("%s...", m.Friendly[0:maxlen])
	}
}
