package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {
<<<<<<< HEAD
	//beginanswer
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*arr = append(*arr, scanner.Text())
	}
	return nil
	//endanswer return nil
}

func ScanToMap(dataMap map[string]string, fileName string) error {
	//beginanswer
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := strings.Split((scanner.Text()), ",")

		dataMap[text[0]] = text[1]
	}
	return nil
	//endanswer return nil
=======
	return nil // TODO: replace this
}

func ScanToMap(dataMap map[string]string, fileName string) error {
	return nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
