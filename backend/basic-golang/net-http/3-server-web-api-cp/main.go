package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// Buat view untuk menampilkan table berdasarkan table ID
// URL: http://localhost:8080/view/table?id=T001

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func viewTable(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	//beginanswer
	var t Table
	id := r.FormValue("id")
	filepath := path.Join("./views", "index.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, d := range data {
		if d.ID == id {
			t = d
			break
		}
	}

	data := map[string]interface{}{
		"title": "Ruangguru Kampus Merdeka",
		"table": t,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/view/table", viewTable)

	fmt.Println("starting web server at http://localhost:8080/")
	// daftarkan server untuk listen di port 8080
	http.ListenAndServe(":8080", nil)
}
