package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
<<<<<<< HEAD
	//beginanswer
	Name         string
	ScoreAverage float64
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

// main function
func main() {
	buff := new(bytes.Buffer)
<<<<<<< HEAD
	//beginanswer

	std := Student{Name: "Rogu", ScoreAverage: 7.8}
	tmp1 := template.New("Template_1")
	tmp1, err := tmp1.Parse("Hello {{.Name}},\nNilai rata-rata kamu {{.ScoreAverage}}\n")
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
