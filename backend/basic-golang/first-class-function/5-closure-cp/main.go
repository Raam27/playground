package main

import "fmt"

func main() {
	//fungsi counter akan menerima (x int) dan mengembalikan fungsi
	//fungsi yang dikembalikan akan melakukan decrement nilai parameter x ketika dipanggil dan
	//mengembalikan nilai parameter x

	counter := func(x int) func() int {
<<<<<<< HEAD
		//beginanswer
		return func() int {
			x--
			return x
		}
		//endanswer
=======
		// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	}
	decrement := counter(5)
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
