package main

import (
	"log"
	"os"
)

// dalam test ini terdapat fungsi os.Remove ya. itu automatis nge remove file yang telah dibuat
// Untuk keperluan testing
func WriteFile(fileName string, fileData string) error {
<<<<<<< HEAD
	//beginanswer

	//membuat file
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}

	//menutup file sebelum fungsi selesai dijalankan
	defer file.Close()

	_, err = file.WriteString(fileData)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
		return err
	}

	return nil
	//endanswer return nil
=======
	return nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
