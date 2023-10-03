package main

import (
	"fmt"
)

func main() {
	// Membuat channel unbuffered
	chInput := make(chan [2]float64)
	chOutput := make(chan string, 2)

	go func() {
		data := <-chInput
		chOutput <- fmt.Sprintf("Hasil Penjumlahan : %f", data[0]+data[1])
		data = <-chInput // Menggunakan variabel yang berbeda
		chOutput <- fmt.Sprintf("Hasil Pengurangan : %f", data[0]-data[1])
	}()

	chInput <- [2]float64{5, 4}
	chInput <- [2]float64{6, 1}

	for i := 0; i < 2; i++ {
		select {
		case result := <-chOutput:
			fmt.Println(result)
		}
	}
}
