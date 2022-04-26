package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {

	//beginanswer
	c := make(chan int)
	//endanswer

	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//kirim 1 ke channel
			//beginanswer
			c <- 1
			//endanswer
		}()
	}
	//mengubah nilai count menggunakan data dari channel

	//beginanswer
	go func() {
		for {
			count += <-c
			wg.Done()
		}
	}()
	//endanswer
	wg.Wait() // menunggu seluruh goroutine selesai berjalan
	output <- count
}
