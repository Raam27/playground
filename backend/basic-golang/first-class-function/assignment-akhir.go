//jangan ditunjukkan ke peserta
package main

import "fmt"

// fungsi ini digunakan untuk menambahkan point
func points(base int) func(x int) int {
	//beginanswer
	return func(points int) int {
		base += points
		return base
		//endanswer
	}
}

func main() {
	toni := points(100) // base value 100
	tono := points(100)
	fmt.Println(toni(2))   // jadi 102
	fmt.Println(toni(3))   // 105
	fmt.Println(toni(100)) // 205
	fmt.Println(tono(2))   // 102
}
