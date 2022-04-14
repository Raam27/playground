package main

import "fmt"

//buat struct Rectangle dengan dua atribut yaitu Width dan Length

<<<<<<< HEAD
//beginanswer
type Rectangle struct {
	Width  int
	Length int
}

//endanswer
=======
// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r) //{10,20}

	r2 := Rectangle{Width: 5, Length: 15}
	fmt.Println(r2) // {5,15}

	fmt.Println("lebar rectangle2:", r2.Width)    //5
	fmt.Println("panjang rectangle2:", r2.Length) //15

}
