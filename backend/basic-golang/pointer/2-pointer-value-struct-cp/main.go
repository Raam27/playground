package main

import "fmt"

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value

<<<<<<< HEAD
//beginanswer
type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) SetWidthPointer(width int) {
	r.Width = width
	fmt.Println("dalam method SetWidthPointer width-nya", r.Width)
}

func (r Rectangle) SetWidthValue(width int) {
	r.Width = width
	fmt.Println("dalam method SetWidthPointer width-nya", r.Width)
}

//endanswer
=======
// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.Width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.Width)
	r.SetWidthValue(70)
	fmt.Println("sesudah melakukan set width dengan value", r.Width)
}
