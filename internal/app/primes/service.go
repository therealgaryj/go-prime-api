package primes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/big"
	"net/http"
	"strconv"
)

var primeCache = make([]bool, 0)

type Service interface {
	GetPrimesList(res http.ResponseWriter, req *http.Request)
}

func NewService() Service {
	return &primesService{}
}

type primesService struct {
}

func (s *primesService) GetPrimesList(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	maxPrime, err := strconv.Atoi(vars[maxPrimesKey])

	if err != nil {
		sendErrorResponse(res, 400, "cannot parse max prime")
		return
	}

	primeNumbers := generatePrimes(maxPrime)

	responseBody := getPrimesResponse{
		Initial: maxPrime,
		Primes:  primeNumbers,
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "application/json")
	res.Header().Add("Foo", "bar")


	jsonResponse, jsonErr := json.Marshal(responseBody)

	if jsonErr != nil {
		sendErrorResponse(res, 500, jsonErr.Error())
		return
	}

	res.Write(jsonResponse)
}

func generatePrimes(maxPrime int) []int {
	primesToReturn := make([]int, 0)

	if maxPrime <= 1 {
		return primesToReturn
	}

	if maxPrime > len(primeCache) + 1 {

		nextPrimes := make([]bool, (maxPrime - len(primeCache)) + 1)
		for x := 0; x < len(nextPrimes); x++ {
			nextPrimes[x] = big.NewInt(int64(len(primeCache) + x)).ProbablyPrime(0)
		}

		primeCache = append(primeCache, nextPrimes...)
	}


	for x := 0; x <= maxPrime; x++ {
		if primeCache[x] {
			primesToReturn = append(primesToReturn, x)
		}
	}

	return primesToReturn
}

func sendErrorResponse(res http.ResponseWriter, status int, description string) {
	res.WriteHeader(status)

	responseBody := errorResponse{Error: description}

	rawBody, jsonErr := json.Marshal(responseBody)

	if jsonErr != nil {
		return
	}

	res.Write(rawBody)
}
