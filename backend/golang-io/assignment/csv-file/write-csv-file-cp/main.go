package main

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteToCSV(filePath string, rows [][]string) error {
	//beginanswer
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	writer.WriteAll(rows)
	//endanswer
	return nil
}

func main() {

	err := WriteToCSV("./users.csv", [][]string{
		{"name", "age", "country"},
		{"levi", "32", "Indonesia"},
		{"jeff", "30", "USA"},
	})

	if err != nil {
		log.Fatal(err)
	}
}
