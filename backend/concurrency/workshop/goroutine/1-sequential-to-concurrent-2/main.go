package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

//Kita cukup menambahkan `go` memanggil fungsi tersebut
func main() {
	go APICallA()
	go APICallB()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("from main function at time", time.Since(start))
}

func APICallA() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("from APICallA at time", time.Since(start))
}

func APICallB() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("from APICallB at time", time.Since(start))
}
