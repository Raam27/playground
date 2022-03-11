package main

import "fmt"

// Kita sudah tau cara membaca array, nah gimana sih kita cara melakukan write
// atau memasukan data ke array ?
func main() {
	var array [10]int
	fmt.Println(array) // output: [0 0 0 0 0 0 0 0 0 0]
	array[0] = 1
	array[1] = 2
	fmt.Println(array) // output: [1 2 0 0 0 0 0 0 0 0]
}
