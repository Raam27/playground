package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan empty interface.
// Buatlah dua data makanan dan minuman yaitu ayam goreng dan cola yang memiliki atribut
// Nama, Jenis, dan Harga.
// Ayam Goreng, Cepat saji, 20000
// Cola, Minuman, 7000

func GetMenu() []map[string]interface{} {
	var menu []map[string]interface{}

<<<<<<< HEAD
	//beginanswer
	ayamGoreng := make(map[string]interface{})
	ayamGoreng["Nama"] = "Ayam Goreng"
	ayamGoreng["Jenis"] = "Cepat saji"
	ayamGoreng["Harga"] = 20000

	menu = append(menu, ayamGoreng)

	cola := make(map[string]interface{})
	cola["Nama"] = "Cola"
	cola["Jenis"] = "Minuman"
	cola["Harga"] = 7000

	menu = append(menu, cola)
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

	return menu
}

func main() {
	menu := GetMenu()

	for _, m := range menu {
		for k, v := range m {
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}
