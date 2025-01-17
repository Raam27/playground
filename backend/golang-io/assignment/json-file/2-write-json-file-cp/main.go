package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//jalankan kode ini di dalam directory yang sama dengan lokasi kode
//hal ini perlu dilakukan agar file yang dicari dapat ditemukan

type student struct {
	Name  string `json:"nama"`
	Score int    `json:"nilai"`
	Class string `json:"kelas"`
}

func main() {
	fileName := "siswa"
	arif := student{Name: "arif", Score: 90, Class: "c"}
	andi := student{Name: "andi", Score: 85, Class: "c"}
	newData := []student{arif, andi}
	writeJSON(fileName, newData)
	reset(fileName) //comment line ini agar file yang dibuat dapat dilihat
}

func writeJSON(fileName string, data []student) error {
<<<<<<< HEAD
	//beginanswer
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		return err
	}
	file, err := openFile(path)
	if err != nil {
		return err
	}
	defer file.Close()
	jsonData, _ := json.Marshal(data)
	ioutil.WriteFile(path, jsonData, 0644)
	return nil
	//endanswer return nil
=======
	return nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func reset(fileName string) {
	os.Remove(fileName + ".json")
}

func (s student) String() string {
	return fmt.Sprintf("name: %v\nscore: %v\nclass: %v\n", s.Name, s.Class, s.Score)
}
