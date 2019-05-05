package internal

import (
	"errors"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"
	dbI "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/persistance"
	log "github.com/micro/go-log"
)

type Service struct {
	db dbI.ContactDB
}

func New(db dbI.ContactDB) *Service {
	return &Service{db}
}

// Call is a single request handler called via client.Call or the generated client code
func (c *Service) Create(contact *model.Contact) error {
	if contact == nil {
		return errors.New("model not found")
	}

	existContact := model.Contact{}
	c.db.GetWithCondition(&existContact, "first_name = ? and email = ?", contact.FirstName, contact.Email)

	if existContact.ID != 0 {
		return errors.New("contact already exist")
	}

	err := contact.BeforeInsert()
	if err != nil {
		return errors.New("can not create contact")
	}

	err = c.db.Create(contact)
	if err != nil {
		return err
	}

	log.Log("Saved succsesfully")
	return nil
}



// Call is a single request handler called via client.Call or the generated client code
func (c *Service) View(id int) (*model.Contact, error) {
	if id == 0 {
		return nil, errors.New("id not found")
	}

	contact := &model.Contact{ID: id}

	err := c.db.View(contact)
	if err != nil {
		return nil, err
	}

	log.Log("Saved successfully")

	return contact, nil
}

// Call is a single request handler called via client.Call or the generated client code
func (c *Service) List() ([]*model.Contact, error) {
	var contacts []*model.Contact

	err := c.db.List(contacts)
	if err != nil {
		return nil, err
	}

	log.Log("Saved successfully")
	return contacts, nil
}
