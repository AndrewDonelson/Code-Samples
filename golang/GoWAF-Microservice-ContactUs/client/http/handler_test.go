package http

import (
	"bytes"
	"errors"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/model"
	"net/url"

	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ContactTestServise struct {
	err error
}

func newTestContactService(err error) *ContactTestServise {
	return &ContactTestServise{err}
}

func (u *ContactTestServise) View(id string) (*model.Contact, error) {
	return nil, u.err
}

func (u *ContactTestServise) List() ([]model.Contact, error) {
	return nil, u.err
}

func (u *ContactTestServise) Create(name, email, category, message string) error {
	return u.err
}

func TestCreateHandler(t *testing.T) {

	data := url.Values{}
	data.Set("message", "Test name")

	tests := []struct {
		description        string
		contactService     *ContactTestServise
		url                string
		method             string
		data               string
		expectedStatusCode int
	}{
		{
			description:        "missing  contact",
			contactService:     newTestContactService(nil),
			url:                "/contact/register",
			method:             http.MethodPost,
			data:               "",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "bad contact",
			contactService:     newTestContactService(nil),
			url:                "/contact/register",
			method:             http.MethodPost,
			data:               "1",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			description:        "error from backend",
			contactService:     newTestContactService(errors.New("error")),
			url:                "/contact/register",
			method:             http.MethodPost,
			data:               data.Encode(),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "all ok",
			contactService:     newTestContactService(nil),
			url:                "/contact/register",
			method:             http.MethodPost,
			data:               data.Encode(),
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {

		contactHandler := NewHandler(tc.contactService, nil)

		r, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer([]byte(tc.data)))
		if err != nil {
			t.Fatal(err)
		}

		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(contactHandler.Create)

		handler.ServeHTTP(w, r)

		resp := w.Result()

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("test: %s.  Unexpected status code %d", tc.description, resp.StatusCode)
		}
	}
}

func TestViewHandler(t *testing.T) {
	badData := url.Values{}
	badData.Set("name", "")

	data := url.Values{}
	data.Set("name", "Test name")

	tests := []struct {
		description        string
		contactService     *ContactTestServise
		url                string
		urlValue           string
		method             string
		data               string
		expectedStatusCode int
	}{
		{
			description:        "missing  id",
			contactService:     newTestContactService(nil),
			url:                "/contact/view/",
			urlValue:           "",
			method:             http.MethodPost,
			data:               "",
			expectedStatusCode: http.StatusBadRequest,
		},

		{
			description:        "error from backend",
			contactService:     newTestContactService(errors.New("error")),
			url:                "/contact/view/",
			urlValue:           "1",
			method:             http.MethodPost,
			data:               data.Encode(),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "pass",
			contactService:     newTestContactService(nil),
			url:                "/contact/edit/",
			urlValue:           "1",
			method:             http.MethodPost,
			data:               data.Encode(),
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {

		contactHandler := NewHandler(tc.contactService, nil)

		r, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer([]byte(tc.data)))
		if err != nil {
			t.Fatal(err)
		}

		r = mux.SetURLVars(r, map[string]string{
			"id": tc.urlValue,
		})
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(contactHandler.View)
		handler.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("test: %s.  Unexpected status code %d", tc.description, resp.StatusCode)
		}
	}
}

func TestListHandler(t *testing.T) {
	badData := url.Values{}
	badData.Set("name", "")

	data := url.Values{}
	data.Set("name", "Test name")

	tests := []struct {
		description        string
		contactService     *ContactTestServise
		url                string
		urlValue           string
		method             string
		data               string
		expectedStatusCode int
	}{
		{
			description:        "error from backend",
			contactService:     newTestContactService(errors.New("error")),
			url:                "/contact/view/",
			urlValue:           "1",
			method:             http.MethodPost,
			data:               data.Encode(),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			description:        "pass",
			contactService:     newTestContactService(nil),
			url:                "/contact/edit/",
			urlValue:           "1",
			method:             http.MethodPost,
			data:               data.Encode(),
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {

		contactHandler := NewHandler(tc.contactService, nil)

		r, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer([]byte(tc.data)))
		if err != nil {
			t.Fatal(err)
		}

		r = mux.SetURLVars(r, map[string]string{
			"id": tc.urlValue,
		})
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(contactHandler.List)
		handler.ServeHTTP(w, r)
		resp := w.Result()

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("test: %s.  Unexpected status code %d", tc.description, resp.StatusCode)
		}
	}
}
