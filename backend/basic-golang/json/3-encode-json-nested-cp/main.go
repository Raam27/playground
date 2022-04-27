package main

import (
	"encoding/json"
	"log"
)

// Dari contoh sebelumnya
// buat JSON string array nested seperti berikut
/*
[
{
		"jenis": "Meja Makan",
		"warna": "Coklat",
		"jumlah": 20,
		"ukuran": {
				"panjang": "50 cm",
				"tinggi": "25 cm"
		}
},
{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
}
]
*/

type Ukuran struct {
<<<<<<< HEAD
	//beginanswer
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
	//endanswer
}

type Meja struct {
	//beginanswer
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	Ukuran Ukuran `json:"ukuran"`
	//endanswer
=======
	// TODO: answer here
}

type Meja struct {
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
<<<<<<< HEAD
	//beginanswer

	mejaJSON, err := json.Marshal(m.MejaMeja)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}

	return string(mejaJSON)

	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func NewMeja(m Items) Items {
	return m
}
