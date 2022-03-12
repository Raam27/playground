package main

import "fmt"

func main() {
	//beginanswer
	decrement := func(x int) func() int {
		return func() int {
			x--
			return x
		}
	}
	//endanswer
	counter := decrement(5)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
