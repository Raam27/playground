package main

import (
	"fmt"
	"math"
)

// Check Point:
// Two Dimention Area
// - Input: Nilai Pertama, Nilai Kedua
// - Output: Result Operation

// Contoh:
// Input:
// - Masukkan nilai pertama: 10
// - Masukkan nilai pertama: 5
// 1: Addition (+)
// 2: Subtraction (-)
// 3: Multiplication (*)
// 4: Division (/)
// - Enter choice: 2
// Output:
// - Subtraction is: 5

func main() {
	//beginanswer
	var choice int = 0
	var result float32
	fmt.Println("1: Rectange Area")
	fmt.Println("2: Rectangular")
	fmt.Println("3: Triangle")
	fmt.Println("4: Circle")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var side float32
		fmt.Printf("Masukkan sisi : ")
		fmt.Scan(&side)
		result = side * side
		fmt.Printf("Luas Persegi adalah : %f", result)
	case 2:
		var length, wide float32
		fmt.Printf("Masukkan panjang : ")
		fmt.Scan(&length)
		fmt.Printf("Masukkan lebar : ")
		fmt.Scan(&wide)
		result = length * wide
		fmt.Printf("Luas Persegi adalah : %f", result)
	case 3:
		var a, t float32
		fmt.Print("Masukkan panjang alas segitga: ")
		fmt.Scanln(&a)
		fmt.Print("Masukkan tinggi segitiga: ")
		fmt.Scanln(&t)
		result = 0.5 * a * t
		fmt.Println("Luas segitiga adalah", result)
	case 4:
		var r, res float64
		const pi float64 = 3.14
		fmt.Printf("Masukkan jari-jari : ")
		fmt.Scan(&r)
		res = pi * (math.Pow(r, 2))
		fmt.Printf("Jadi luasnya adalah : %f", res)
	default:
		fmt.Println("Invalid value")
	}
	//endanswer
}
