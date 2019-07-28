package common

import (
	"math"
	"math/rand"
	"time"
)

// RandomNumber generates unique random num
func RandomNumber(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(max)
}

// RandomMinNumber generates unique random number
// that should be higher than min and lower than max
func RandomMinNumber(min, max int) int {
	if min > max {
		panic("min number cannot be higher than max")
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	num := r.Intn(max)
	if num < min {
		num = int(math.Min(float64(num+min), float64(max)))
	}

	return num
}