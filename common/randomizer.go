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
