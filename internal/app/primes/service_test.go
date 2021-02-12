package primes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimesService_generatePrimes_10(t *testing.T) {

	primeCache = make([]bool, 0)

	expected := []int{2, 3, 5, 7}
	actual := generatePrimes(10)

	assert.Equal(t, expected, actual)
}

func TestPrimesService_generatePrimes_repeated(t *testing.T) {

	primeCache = make([]bool, 0)

	expected := []int{2, 3, 5, 7}
	actual := generatePrimes(10)
	actual = generatePrimes(10)

	assert.Equal(t, expected, actual)
}

func TestPrimesService_generatePrimes_100(t *testing.T) {

	primeCache = make([]bool, 0)

	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	actual := generatePrimes(100)

	assert.Equal(t, expected, actual)
}