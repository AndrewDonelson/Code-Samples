package handler

import (
	"context"
	"errors"
	"testing"

	models "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"
	dbI "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/persistance"
	contact "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/proto"
)

type testContactDB struct {
	Err error
	id  int
	contacts  []models.Contact
}

func newTestContactDB(err error, id int, contacts []models.Contact) *testContactDB {
	return &testContactDB{err, id, contacts}
}

func (gpg *testContactDB) Create(model *models.Contact) error {
	model.ID = gpg.id
	return gpg.Err
}

func (gpg *testContactDB) View(model *models.Contact) error {
	model.ID = gpg.id
	return gpg.Err
}

func (gpg *testContactDB) List(model []models.Contact) error {
	model = gpg.contacts
	return gpg.Err
}

func (gpg *testContactDB)GetWithCondition(model *models.Contact, condition interface{}, args ...interface{}) error{
	model.ID = gpg.id
	return gpg.Err
}

func TestContact_Create(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		db dbI.ContactDB
	}

	type args struct {
		ctx context.Context
		req *contact.CreateRequest
		rsp *contact.CreateResponse
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"All ok", fields{newTestContactDB(nil, 0, nil)}, args{ctx, createContactRequest("name", "test@mail.com", "test", "test category"), createContactResponce("")}, false},
		{"bad email", fields{newTestContactDB(nil, 0, nil)}, args{ctx, createContactRequest("name", "test", "test", "test category"), createContactResponce("")}, true},
		{"error from db", fields{newTestContactDB(errors.New("error"), 0, nil)}, args{ctx, createContactRequest("name", "test@mail.com", "test", "test category"), createContactResponce("")}, true},
		{"contact already exist", fields{newTestContactDB(nil, 1, nil)}, args{ctx, createContactRequest("name", "test@mail.com", "test", "test category"), createContactResponce("")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contact{
				db: tt.fields.db,
			}
			if err := c.Create(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Contact.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createContactRequest(name, email, message, category string) *contact.CreateRequest {
	req := &contact.CreateRequest{Name:name, Email:email, Message:message, Category:category}
	return req
}
func createContactResponce(errMsg string) *contact.CreateResponse {
	rsp := &contact.CreateResponse{ErrMsg: errMsg}
	return rsp
}


func TestContact_View(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		db dbI.ContactDB
	}

	type args struct {
		ctx context.Context
		req *contact.ViewRequest
		rsp *contact.ViewResponse
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"All ok", fields{newTestContactDB(nil, 0, nil)}, args{ctx, viewContactRequest("1"), viewContactResponce()}, false},
		{"error from db", fields{newTestContactDB(errors.New("error"), 0, nil)}, args{ctx, viewContactRequest("1"), viewContactResponce()}, true},
		{"empty id", fields{newTestContactDB(nil, 0, nil)}, args{ctx, viewContactRequest(""), viewContactResponce()}, true},
		{"bad id", fields{newTestContactDB(nil, 0, nil)}, args{ctx, viewContactRequest("x"), viewContactResponce()}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contact{
				db: tt.fields.db,
			}
			if err := c.View(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Contact.View() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func viewContactRequest(id string) *contact.ViewRequest {
	req := &contact.ViewRequest{ContactID: id}
	return req
}
func viewContactResponce() *contact.ViewResponse {
	return &contact.ViewResponse{}
}


func TestContact_List(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		db dbI.ContactDB
	}
	type args struct {
		ctx context.Context
		req *contact.Empty
		rsp *contact.ListResponse
	}

	contacts := []models.Contact{
		{ID:1},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"not find models", fields{newTestContactDB(nil, 1, contacts)}, args{ctx, nil, listContactResponce("")}, true},
		{"error drom db", fields{newTestContactDB(errors.New("error"), 1, contacts)}, args{ctx, nil, listContactResponce("")}, true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contact{
				db: tt.fields.db,
			}
			if err := c.List(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Contact.List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func listContactResponce(errMsg string) *contact.ListResponse {
	contacts := []*contact.Contact{
		{ContactID:"1"},
	}

	rsp := &contact.ListResponse{Contacts: contacts, ErrMsg: errMsg}
	return rsp
}
