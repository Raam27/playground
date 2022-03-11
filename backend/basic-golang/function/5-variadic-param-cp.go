package main

import "fmt"

func main() {
	printWord("halo")
	printWord("halo", "selamat siang")
	printWord("halo", "andi", "dan", "bagus")
	printWord("mencoba", "variadic", "param", "pada", "go")
}

//beginanswer
func printWord(words ...string) {
	for _, word := range words {
		fmt.Println(word)
	}
}

//endanswer
