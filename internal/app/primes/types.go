package primes

type getPrimesResponse struct {
	Initial int `json:"initial"`
	Primes []int `json:"primes"`
}

type errorResponse struct {
	Error string `json:"error"`
}
