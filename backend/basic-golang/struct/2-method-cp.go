package main

import "fmt"

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan area
// GetPerimeter() digunakan untuk menampilkan keliling

//beginanswer
type Rectangle struct {
	Width  int
	Length int
}

func (r Rectangle) GetArea() {
	fmt.Println("luas persegi panjang : ", r.Width*r.Length)
}

func (r Rectangle) GetPerimeter() {
	fmt.Println("keliling persegi panjang : ", 2*(r.Width+r.Length))
}

//endanswer
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
