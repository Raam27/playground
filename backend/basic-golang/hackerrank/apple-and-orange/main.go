package main

import (
	"fmt"
)

// Apple and Orange
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/apple-and-orange/problem

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
<<<<<<< HEAD
	//beginanswer
	count := func(base int32, fruit []int32) {
		var c int32
		for _, d := range fruit {
			if base+d >= s && base+d <= t {
				c++
			}
		}
		fmt.Println(c)
	}
	count(a, apples)
	count(b, oranges)
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func main() {
	var apples = []int32{3, 2, -2}
	var oranges = []int32{2, 1, 5, -6}

	countApplesAndOranges(7, 11, 5, 15, apples, oranges)
}
