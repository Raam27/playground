package main

import "fmt"

//pakai * untuk melakukan perkalian
//misal number1 * number2
func main() {
	result := multiply(3, 3)
	fmt.Println(result)
	fmt.Println(multiply(5, 5))
}

//beginanswer
func multiply(number1, number2 int) int {
	result := number1 * number2
	return result
}

//endanswer