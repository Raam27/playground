package main

import "fmt"

//fungsi square akan mengembalikan nilai pangkat 2
//dari dua parameter yang diterima
//contoh
//parameter yang diterima (4,3)
//maka akan mengembalikan 16 dan 9
func main() {

	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

//gunakan * untuk melakukan perkalian
<<<<<<< HEAD
//beginanswer
func square(angka1, angka2 int) (int, int) {
	result1 := angka1 * angka1
	result2 := angka2 * angka2
	return result1, result2
}

//endanswer
=======
// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
