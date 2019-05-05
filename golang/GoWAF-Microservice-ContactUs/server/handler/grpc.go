package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/internal"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"

	"github.com/micro/go-log"
	"strconv"

	contact "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/proto"
)

type Contact struct {
	srv *internal.Service
}

func New(db *internal.Service) *Contact {
	return &Contact{srv: db}
}

// Call is a single request handler called via client.Call or the generated client code
func (c *Contact) Create(ctx context.Context, req *contact.CreateRequest, rsp *contact.CreateResponse) error {
	contact := contactRegisterRequest(req)
	err := c.srv.Create(contact)
	if err != nil {
		rsp.ErrMsg = err.Error()
		return err
	}

	log.Log("Saved successfully")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (c *Contact) View(ctx context.Context, req *contact.ViewRequest, rsp *contact.ViewResponse) error {
	if req.ContactID == "" {
		return errors.New("could not get id")
	}
	id, _ := strconv.Atoi(req.ContactID)
	contact, err := c.srv.View(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rsp.Contact = toContact(contact)
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (c *Contact) List(ctx context.Context, req *contact.Empty, rsp *contact.ListResponse) error {
	contacts, err := c.srv.List()
	if err != nil {
		rsp.ErrMsg = err.Error()
		return err
	}

	rsp.Contacts = toContactList(contacts)
	return nil
}

//convert model.ContactInfo to contact.Contact
func toContact(c *model.Contact) *contact.Contact {
	var contact contact.Contact

	contact.ContactID = strconv.Itoa(c.ID)
	contact.Email = c.Email
	contact.Message = c.Message
	contact.Name = c.FirstName
	contact.Category = c.Category
	contact.CreatedAt = c.CreatedAt.Format("01-02-2006 15:04:05")
	return &contact
}

//convert []model.ContactInf to []*contact.Contact
func toContactList(cc []*model.Contact) []*contact.Contact {
	var contactList []*contact.Contact

	for _, c := range cc {
		var contact contact.Contact

		contact.ContactID = strconv.Itoa(c.ID)
		contact.Email = c.Email
		contact.Name = c.FirstName
		contact.Message = c.Message
		contact.Category = c.Category
		contact.CreatedAt = c.CreatedAt.Format("01-02-2006 15:04:05")
		contactList = append(contactList, &contact)
	}

	return contactList
}

func contactRegisterRequest(req *contact.CreateRequest) *model.Contact {
	var contact model.Contact
	contact.FirstName = req.Name
	contact.Email = req.Email
	contact.Category = req.Category
	contact.Message = req.Message
	return &contact
}
