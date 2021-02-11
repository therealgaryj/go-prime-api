package main

import (
	"github.com/therealgaryj/go-prime-api/internal/app/primes"
	"net/http"
)

func main() {
	r := primes.Initialize()
	http.ListenAndServe(":8080", r)
}
