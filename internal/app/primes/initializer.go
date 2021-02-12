package primes

import (
	"github.com/gorilla/mux"
	"net/http"
)

const maxPrimesKey = "maxPrime"

func Initialize() *mux.Router {

	primesService := NewService()

	r := mux.NewRouter()

	r.Path("/primes/{" + maxPrimesKey + ":.+}").
		Methods(http.MethodGet).
		HandlerFunc(primesService.GetPrimesList)

	return r
}
