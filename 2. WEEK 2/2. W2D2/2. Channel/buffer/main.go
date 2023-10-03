package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	// Go routine Pengirim
	go func() {
		fmt.Println("Mulai Mengirim Pesan...")
		ch <- "Hello dari unbuffered channel!"
		fmt.Println("pesan terkirim")
	}()

	time.Sleep(2 * time.Second)
	// Go Routine Penerima
	go func() {
		message := <-ch
		fmt.Println("Pesan diterima:", message)
	}()

	time.Sleep(3 * time.Second)
}
