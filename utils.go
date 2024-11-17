package main

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// 1,00,000 --> 8,99,999 + 1,00,000 --> 18,99,999
// 10,00,000 99,99,999
// generate random integer ID of specified digits
func IdGenerator(digits int) (int, error) {

	if digits <= 0 {
		return -1, errors.New("digits limit is not valid. It should be greater than 0")
	}
	lowerBound := int(math.Pow10(digits - 1))
	upperBound := int(math.Pow10(digits))

	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(upperBound-lowerBound) + lowerBound, nil
}
