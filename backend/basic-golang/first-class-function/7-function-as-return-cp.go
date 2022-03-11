package main

import "fmt"

func main() {
	// fungsi ini mengembalikan
	// func(x, y int) int {
	// 	return x * y
	// }
	//beginanswer
	getAreaFunc := func() func(int, int) int {
		return func(x, y int) int {
			return x * y
		}
	}
	//endanswer
	areaF := getAreaFunc()
	res := areaF(8, 4)
	fmt.Println(res)
}
