package common

import (
	"math/rand"
	"time"
)

// RandomNumber generates unique random num
func RandomNumber(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(max)
}

func RandomMinNumber(min, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if num < min {
		num += min
	}

	return num
}
