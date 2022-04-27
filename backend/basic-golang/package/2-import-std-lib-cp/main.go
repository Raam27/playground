package main

import (
<<<<<<< HEAD
	//beginanswer
	"fmt"
	"time"
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
)

// Dari contoh yang telah diberikan dan eksplorasi yang dilakukan dari standard library golang, kamu dapat mencoba untuk mengimport salah satu package pada golang.
// Cobalah untuk mengimport package time dan lakukan perhitungan perbedaan hari antara dua waktu.

func CountDays(start, end time.Time) int {
<<<<<<< HEAD
	//beginanswer
	days := end.Sub(start).Hours() / 24
	return int(days)
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func main() {
	start := time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC)
	dayDifference := CountDays(start, end)
	fmt.Println(dayDifference)
}
