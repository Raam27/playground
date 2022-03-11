package main

import "fmt"

func main() {
	// bisa disimpan dalam variabel
	result := sum(3, 3)
	fmt.Println(result)
	fmt.Println(sum(5, 5))
}

func sum(number1, number2 int) int {
	result := number1 + number2
	return result
}
