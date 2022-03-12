package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Eat(food string) {
	fmt.Printf("%s makan %s\n", p.Name, food)
}

func (p Person) Sleep() {
	fmt.Printf("%s tidur\n", p.Name)
}

func main() {
	var p Person
	p.Name = "putra"
	p.Age = 20
	fmt.Println(p)
	p.Eat("nasi goreng")
	p.Sleep()
}
