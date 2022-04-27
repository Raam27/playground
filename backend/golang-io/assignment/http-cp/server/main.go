package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		//beginanswer
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
		//endanswer
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		//beginanswer
		data := r.URL.Query().Get("data")
		w.Write([]byte(data))
		//endanswer
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		//beginanswer
		a := r.URL.Query().Get("a")
		aInt, err := strconv.Atoi(a)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b := r.URL.Query().Get("b")
		bInt, err := strconv.Atoi(b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sum := aInt + bInt
		w.Write([]byte(fmt.Sprint(sum)))
		//endanswer
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		//beginanswer
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, world!"}`))
		//endanswer
=======
		// TODO: answer here
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
