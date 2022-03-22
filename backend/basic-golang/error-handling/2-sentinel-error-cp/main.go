package main

import (
	// import
	"errors"
)

// Dari contoh yang diberikan, buatlah sentinel error tambahan untuk error handling pada function FindData
// Kalian dapat menambahkan sentinel error untuk mengecek apakah umurnya valid (misal kurang dari 0)

var ErrDataNotFound = errors.New("error data not found")

//beginanswer
var ErrInvalidAge = errors.New("error age is invalid, less than 0")

//endanswer
