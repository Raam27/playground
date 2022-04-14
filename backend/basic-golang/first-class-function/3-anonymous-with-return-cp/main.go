package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
<<<<<<< HEAD
		//beginanswer
		return "Selamat sore"
		//endanswer
=======
		// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	}()

	fmt.Println(goodAfternoon)
}
