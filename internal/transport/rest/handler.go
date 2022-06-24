package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	param         = "name"
	httpHeaderA   = "a"
	httpHeaderB   = "b"
	httpHeaderSum = "a+b"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(h.notFoundHandler)

	r.HandleFunc("/name/{"+param+"}", h.getName).Methods(http.MethodGet)
	r.HandleFunc("/bad", h.badRequest).Methods(http.MethodGet)
	r.HandleFunc("/data", h.getDataFromBody).Methods(http.MethodPost)
	r.HandleFunc("/headers", h.getDataFromHeaders).Methods(http.MethodPost)
	return r
}

func (h *Handler) getName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, existed := vars[param]
	if !existed {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := []byte(fmt.Sprintf("Hello, %s!\n", name))
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write(response)
}

func (h *Handler) badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *Handler) getDataFromBody(w http.ResponseWriter, r *http.Request) {
	message, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := []byte(fmt.Sprintf("I got message:\n%s", string(message)))
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write(response)
}

func (h *Handler) getDataFromHeaders(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get(httpHeaderA))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(r.Header.Get(httpHeaderB))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sum := strconv.Itoa(a + b)
	w.Header().Add(httpHeaderSum, sum)
}

func (h *Handler) notFoundHandler(w http.ResponseWriter, r *http.Request) {

}
