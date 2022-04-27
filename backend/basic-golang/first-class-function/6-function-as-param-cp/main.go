package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
<<<<<<< HEAD
	//beginanswer
	printString := func(f func() string) {
		result := f()
		fmt.Println(result)
	}

	goodNight := func() string {
		return "good night friends"
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

	printString(goodNight)

}
