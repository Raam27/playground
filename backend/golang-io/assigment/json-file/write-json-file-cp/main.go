package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Name    string
	Age     int
	Country string
}
type Users []User

func WriteToJson(filePath string, users Users) error {
	//beginanswer
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	json.NewEncoder(file).Encode(users)

	//endanswer
	return nil
}

func main() {
	err := WriteToJson("./users.json", Users{
		{
			Name:    "levi",
			Age:     32,
			Country: "Indonesia",
		},
		{
			Name:    "jeff",
			Age:     30,
			Country: "USA",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
