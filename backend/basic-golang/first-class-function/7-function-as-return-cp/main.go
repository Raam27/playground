package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
<<<<<<< HEAD
	//beginanswer
	getAreaFunc := func() func(int, int) int {
		return func(x, y int) int {
			return x * y
		}
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	areaF := getAreaFunc()
	res := areaF(3, 4) // 12
	fmt.Println(res)
}
