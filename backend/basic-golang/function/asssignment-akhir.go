//jangan ditunjukkan ke peserta
package main

import "fmt"

func main() {
	sumResult, subtractResult, multiplyResult, divideResult := calculate(4, 4)
	fmt.Println("hasil penjumlahan", sumResult)
	fmt.Println("hasil pengurangan", subtractResult)
	fmt.Println("hasil perkalian", multiplyResult)
	fmt.Println("hasil pembagian", divideResult)
	fmt.Println(calculate(5, 5))
}

func calculate(Number1, Number2 int) (sumResult, subtractResult, multiplyResult, divideResult int) {
	//beginanswer
	sumResult = Number1 + Number2
	subtractResult = Number1 - Number2
	multiplyResult = Number1 * Number2
	divideResult = Number1 / Number2
	//endanswer
	return
}
