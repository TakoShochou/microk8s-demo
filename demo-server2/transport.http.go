package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPHandler() http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/ready").Handler()
	return r
}
