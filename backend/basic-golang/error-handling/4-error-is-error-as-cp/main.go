package main

import (
	"errors"
	"fmt"
)

// Dari contoh yang telah diberikan, coba kamu gunakan errors.Is untuk mengecek jenis error tertentu.
// Modifikasikan pada kode dari wrapping error
var ErrDataNotFound = errors.New("error data not found")

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, ErrDataNotFound
	}

	return data[name], nil
}

func IsEligibleToVaccine(data map[string]int, name string) (bool, error) {
	age, err := GetAge(data, name)
	if err != nil {
<<<<<<< HEAD
		//beginanswer
		return false, fmt.Errorf("error in IsEligibleToVaccine, error = %w", err)
		//endanswer
=======
		// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	}
	if age < 15 {
		return false, nil
	}

	return true, nil
}

func main() {
	data := map[string]int{
		"Roger":  60,
		"Kamala": 5,
		"Peter":  20,
	}
	_, err := IsEligibleToVaccine(data, "Tony")
	if err != nil {
		// Cek apakah err merupakan jenis error ErrDataNotFound
<<<<<<< HEAD
		//beginanswer
		if errors.Is(err, ErrDataNotFound) {
			fmt.Println("err adalah ErrDataNotFound")
		}
		//endanswer
=======
		// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	}
}
