package main

import "fmt"

// pada cekpoin ini kalian akan mencoba untuk menghapus data dengan tipe []map[string]string
// gabungan slice dan map.

func main() {
	var namaUmur = []map[string]string{
		{"name": "Socrates", "gender": "male"},
		{"name": "Plato", "gender": "male"},
		{"name": "Ada", "gender": "female"},
		{"name": "Leonhard Euler", "gender": "female"},
		{"name": "Blaise Pascal", "gender": "male"},
	}

	// terdapat kesalahan pada data gender tersebut dapatkan kalian memperbaiki nya ?
<<<<<<< HEAD
	//beginanswer
	namaUmur[3]["gender"] = "male"
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	for _, val := range namaUmur {
		fmt.Println(val["name"], " ", val["gender"])
	}

	// Nah coba saatnya kalian menghapuskan key "gender" pada setiap data
	// delete data if key is equal "gender"

	for _, val := range namaUmur {
		fmt.Println(val)
	}
	// Output sebelum dihapus
	/*
		map[gender:male name:Socrates]
		map[gender:male name:Plato]
		map[gender:female name:Ada]
		map[gender:male name:Leonhard Euler]
		map[gender:male name:Blaise Pascal]
	*/

<<<<<<< HEAD
	//beginanswer
	for _, val := range namaUmur {
		delete(val, "gender")
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9

	// Output setelah dihapus
	/*
		map[name:Socrates]
		map[name:Plato]
		map[name:Ada]
		map[name:Leonhard Euler]
		map[name:Blaise Pascal]
	*/
	for _, val := range namaUmur {
		fmt.Println(val)
	}
}
