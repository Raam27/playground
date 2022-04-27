package main

import "fmt"

// Check Point:
// Start Pattern
// - Input: Size
// - Output: Start Pattern

// Contoh:
// - Input: Size: 10
// - Output:
//           *
//          **
//         ***
//        ****
//       *****
//      ******
//     *******
//    ********
//   *********
//  **********

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

<<<<<<< HEAD
	//beginanswer
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
