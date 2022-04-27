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

<<<<<<< HEAD
//beginanswer
type user struct {
	Name   string `json:"name"`
	School string `json:"age"`
	Class  string `json:"email"`
	Phone  string `json:"role"`
	Score  int    `json:"status"`
}

//endanswer
=======
// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

func (s user) String() string {
	return fmt.Sprintf("name: %v\nschool: %v\nclass: %v\nphone: %v\nscore: %v\n", s.Name, s.School, s.Class, s.Phone, s.Score)
}

func main() {
	fileName := "data/read"
	json, err := readJSON(fileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, user := range json {
		fmt.Println(user)
	}
}

func readJSON(fileName string) ([]user, error) {
<<<<<<< HEAD
	//beginanswer
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		return nil, err
	}
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	users := []user{}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &users)
	return users, nil
	//endanswer []user{}, nil
=======
	[]user{}, nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
