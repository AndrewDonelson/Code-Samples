package grpc

import (
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/model"
)

type ContactMicroservice interface {
	Create (name, email, category, message string) error
	View(id string) (*model.Contact, error)
	List() ([]model.Contact, error)
}
