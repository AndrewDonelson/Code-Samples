package grpc

import (
	"fmt"
	proto "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/grpc/proto"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/model"
	"strconv"
	"time"

	"context"
)

type ContactMicroGrpc struct {
	service proto.ContactMicroService
}

//Return new ContactMicroGrpc
func New(us proto.ContactMicroService) *ContactMicroGrpc {
	return &ContactMicroGrpc{service: us}
}

// Login displays login page for GET and processes on POST
func (h *ContactMicroGrpc) Create(name, email, message, category string) error {

	// call the backend service
	_, err := h.service.Create(context.TODO(), &proto.CreateRequest{
		Name:     name,
		Email:    email,
		Message:  message,
		Category: category,
	})

	return err
}

// Login displays login page for GET and processes on POST
func (h *ContactMicroGrpc) View(id string) (*model.Contact, error) {

	// call the backend service
	contact, err := h.service.View(context.TODO(), &proto.ViewRequest{
		ContactID: id,
	})

	if err != nil {
		return nil, err
	}

	return toModelContact(contact.Contact), nil
}

// Login displays login page for GET and processes on POST
func (h *ContactMicroGrpc) List() ([]model.Contact, error) {

	// call the backend service
	rsp, err := h.service.List(context.TODO(), &proto.Empty{})

	if err != nil {
		return nil, err
	}

	return toModeslContacts(rsp.Contacts), err
}

//convert model.ContactInfo to contact.Contact
func toModelContact(c *proto.Contact) *model.Contact {
	var contact = new(model.Contact)
	idInt, err := strconv.Atoi(c.ContactID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	contact.ID = idInt
	contact.FirstName = c.Name
	contact.Email = c.Email

	contact.Email = c.Email

	t, err := time.Parse("01-02-2006 15:04:05", c.CreatedAt)
	if err != nil {
		fmt.Println(err)
	}

	contact.CreatedAt = t
	return contact
}

//convert []model.ContactInf to []*contact.Contact
func toModeslContacts(uu []*proto.Contact) []model.Contact {
	var contacts []model.Contact

	for _, u := range uu {
		var contact model.Contact

		idInt, err := strconv.Atoi(u.ContactID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		contact.ID = idInt
		contact.Email = u.Email
		contact.FirstName = u.Name

		t, err := time.Parse("01-02-2006 15:04:05", u.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}

		contact.CreatedAt = t
		contacts = append(contacts, contact)
	}
	return contacts
}
