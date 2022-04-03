package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}
type Users []User

func ReadUsers(filePath string) (Users, error) {
	//beginanswer
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	users := make(Users, 0)
	json.NewDecoder(file).Decode(&users)

	return users, nil
	//endanswer panic
}

func main() {
	users, err := ReadUsers("./users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", users)
}
