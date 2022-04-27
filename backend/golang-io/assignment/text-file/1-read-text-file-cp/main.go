package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
<<<<<<< HEAD
	//beginanswer
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return FileData{}, err
	}

	return FileData{
		Name: name,
		Size: len(data),
		Data: data,
	}, nil
	//endanswer return FileData{}, nil
=======
	return FileData{}, nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
