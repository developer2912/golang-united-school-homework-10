package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	NAME = "name"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/name/{"+NAME+":[a-zA-Z9-0_]+}", h.getName).Methods(http.MethodGet)
	r.HandleFunc("/bad", h.badRequest).Methods(http.MethodGet)
	r.HandleFunc("/data", h.getDataFromBody).Methods(http.MethodPost)
	r.HandleFunc("/headers", h.getDataFromHeaders).Methods(http.MethodPost)
	return r
}

func (h *Handler) getName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, existed := vars[NAME]
	if !existed {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := []byte(fmt.Sprintf("Hello, %s!\n", name))
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *Handler) getDataFromBody(w http.ResponseWriter, r *http.Request) {
	message, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := []byte(fmt.Sprintf("I got message:\n%s", string(message)))
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) getDataFromHeaders(w http.ResponseWriter, r *http.Request) {
	a, b := r.Header.Get("a"), r.Header.Get("b")
	w.Header().Add("a+b", a+b)
}
