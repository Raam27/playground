package main

import (
	"encoding/json"
	"log"
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	//beginanswer
	Jenis  string `json:"jenis"`
	Warna  string `json:"color"`
	Jumlah int    `json:"jumlah"`
	//endanswer
}

func (m Meja) EncodeJSON() string {
	//beginanswer

	mejaJSON, err := json.Marshal(m)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}

	return string(mejaJSON)

	//endanswer
}

func NewMeja(m Meja) Meja {
	return m
}
