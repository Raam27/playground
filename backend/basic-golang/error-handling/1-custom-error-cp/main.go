package main

import (
	"errors"
	"fmt"
)

// Dari contoh yang telah diberikan, kalian dapat mencoba membuat custom error baru dengan atribut message dan errCode.
// Misalnya adalah error untuk validasi data umur kurang dari 0.

//beginanswer
type ErrorInvalidData struct {
	message string
	errCode int32
}

func (e *ErrorInvalidData) Error() string {
	return fmt.Sprintf("error %d: %s", e.errCode, e.message)
}

//endanswer

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, errors.New("Data not found")
	}

	if data[name] < 0 {
		// Isilah baris ini dengan return 0 dan custom error yang telah dibuat
		//beginanswer
		return 0, &ErrorInvalidData{
			message: "error invalid data",
			errCode: 500,
		}
		//endanswer
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
