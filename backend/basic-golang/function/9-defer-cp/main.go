package main

import "fmt"

func main() {
	order("teh")
	order("kopi")
}

func order(drink string) {
	// mengucapkan terima kasih di akhir dengan menggunakan defer
<<<<<<< HEAD
	//beginanswer
	defer fmt.Println("terima kasih, silahkan datang kembali")
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
	fmt.Println("kami sedang ada diskon untuk pembelian kopi")
	fmt.Println("pesanan anda:", drink)
	if drink == "kopi" {
		fmt.Println("okee, karena diskon totalnya jadi 4000")
	}
	fmt.Println("mohon ditunggu")
}
