package http

import (
	ctx "github.com/NlaakStudiosLLC/Company-Website/webapp/context"
	communication "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/grpc"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/gorilla/mux"
)

// Decoder is use to decode the schema
var Decoder = schema.NewDecoder()

type ContactHandler struct {
	service communication.ContactMicroservice
	Ctx     *ctx.Context
}

//Return new Service
func NewHandler(us communication.ContactMicroservice, Ctx *ctx.Context) *ContactHandler {
	return &ContactHandler{service: us, Ctx: Ctx}
}

//Register - register new user
func (h *ContactHandler) Create(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request into ContactInfo{}
	contactIncome := model.Contact{}

	if err := decode(r, &contactIncome); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if contactIncome.Message == "" {
		http.Error(w, "missing contact", http.StatusBadRequest)
		return
	}

	// call the backend service
	err := h.service.Create(contactIncome.FirstName, contactIncome.Email, contactIncome.Message, contactIncome.Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}


func (h *ContactHandler) View(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	_, err := h.service.View(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (h *ContactHandler) List(w http.ResponseWriter, r *http.Request) {
	_, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func decode(req *http.Request, model interface{}) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	if err := Decoder.Decode(model, req.PostForm); err != nil {
		return err
	}

	return nil
}
