package main

import (
	// import
	"errors"
	"fmt"
)

// Dari contoh yang diberikan, buatlah sentinel error tambahan untuk error handling pada function FindData
// Kalian dapat menambahkan sentinel error untuk mengecek apakah umurnya valid (misal kurang dari 0) dengan nama ErrInvalidAge

var ErrDataNotFound = errors.New("error data not found")

<<<<<<< HEAD
//beginanswer
var ErrInvalidAge = errors.New("error age is invalid, less than 0")

//endanswer
=======
// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, ErrDataNotFound
	}

	if data[name] < 0 {
		// Isilah baris ini dengan return 0 dan sentinel error ErrInvalidAge
<<<<<<< HEAD
		//beginanswer
		return 0, ErrInvalidAge
		//endanswer
=======
		// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	}

	return data[name], nil
}

func main() {
	peopleAge := map[string]int{
		"John": 20,
		"Raz":  8,
		"Tony": -1,
	}

	_, err := GetAge(peopleAge, "Tony")
	if err != nil {
		fmt.Println(err.Error())
	}
}
