package main

import (
	"encoding/json"
	"log"
)

// Buat string JSON dengan hasil seperti berikut
// [{"jenis":"Meja Lipat","warna":"Coklat","jumlah":40,"deskripsi":"meja untuk belajar"},{"jenis":"Meja Hijau","warna":"Hijau","jumlah":10,"deskripsi":"meja untuk pengadilan"}]

type Meja struct {
<<<<<<< HEAD
	//beginanswer
	Jenis     string `json:"jenis"`
	Warna     string `json:"warna"`
	Jumlah    int    `json:"jumlah"`
	Deskripsi string `json:"deskripsi"`
	//endanswer
=======
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
