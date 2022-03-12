package main

import "fmt"

func main() {
	//fungsi yang mengembalikan nilai pangkat dua dari parameter yang diterima
	//beginanswer
	square := func(angka int) int {
		return angka * angka
	}(5)
	//endanswer
	fmt.Println(square)
}
