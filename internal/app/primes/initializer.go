package primes

import (
	"github.com/gorilla/mux"
	"net/http"
)

const maxPrimesKey = "maxPrime"

func Initialize() *mux.Router {

	notesService := NewService()

	r := mux.NewRouter()

	r.Path("/primes/{" + maxPrimesKey + ":[0-9]+}").
		Methods(http.MethodGet).
		HandlerFunc(notesService.GetPrimesList)

	return r
}
