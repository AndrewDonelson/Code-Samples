package persistance

import 	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"

// TagDB defines the tag db related methods
type ContactDB interface {
	Create(model *model.Contact) error
	View(model *model.Contact) error
	List(model []*model.Contact) error
	GetWithCondition(model *model.Contact, condition interface{}, args ...interface{}) error
}