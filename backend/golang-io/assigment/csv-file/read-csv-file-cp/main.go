package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Name    string
	Age     int
	Country string
}
type Users []User

func ReadUsers(filePath string) (Users, error) {
	//beginanswer
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	users := make(Users, 0)
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		age, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, err
		}
		user := User{
			Name:    row[0],
			Age:     age,
			Country: row[2],
		}
		users = append(users, user)
	}

	return users, nil
	//endanswer
}

func main() {
	users, err := ReadUsers("./users.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", users)
}
