package mygo

import "math/rand"

// ReadNumber create random number
func ReadNumber() int {
	rnr := 10
	return rand.Intn(rnr)
}
