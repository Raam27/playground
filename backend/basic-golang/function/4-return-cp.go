package main

import "fmt"

func main() {
	//pakai * untuk melakukan perkalian
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
