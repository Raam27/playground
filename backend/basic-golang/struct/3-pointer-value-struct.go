package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) BirthdayPointer() {
	p.Age++
	fmt.Println("dalam method BirhtdayPointer umurnya", p.Age)
}

func (p Person) BirthdayValue() {
	p.Age++
	fmt.Println("dalam method BirthdayValue umurnya", p.Age)

}

func main() {
	var p Person
	p.Name = "putra"
	p.Age = 20
	fmt.Println("sebelum ulang tahun", p.Age)
	p.BirthdayPointer()
	fmt.Println("setelah memanggil BirthdayPointer", p.Age)

	p.BirthdayValue()
	fmt.Println("setelah memanggil BirthdayValue", p.Age)
}
