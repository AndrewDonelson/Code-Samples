package http

import (
	"github.com/gorilla/mux"

	"net/http"
)

//Init- init routs
func InitRouters (c *ContactHandler) *mux.Router {
	r := mux.NewRouter()

	csr := r.PathPrefix("/contact").Subrouter()
	csr.HandleFunc("/create", c.Create).Methods(http.MethodPost)
	csr.HandleFunc("/list", c.List).Methods(http.MethodGet)
	csr.HandleFunc("/{id}", c.View).Methods(http.MethodGet)

	return r
}
