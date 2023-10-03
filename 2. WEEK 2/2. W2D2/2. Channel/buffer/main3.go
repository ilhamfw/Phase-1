package main

import (
	"fmt"
)

func main() {
	// Membuat channel unbuffered
	chInput := make(chan [2]float64)
	chAddOutput := make(chan float64)
	chSubOutput := make(chan float64)

	go func() {
		data := <-chInput
		chAddOutput <- data[0] + data[1]
	}()

	go func() {
		data := <-chInput
		chSubOutput <- data[0] - data[1]
	}()

	chInput <- [2]float64{5, 4} // Perbaiki penggunaan tipe data dengan {} bukan ()
	chInput <- [2]float64{6, 1} // Perbaiki penggunaan tipe data dengan {} bukan ()

	fmt.Println("Hasil Penjumlahan :", <-chAddOutput)
	fmt.Println("Hasil Pengurangan :", <-chSubOutput)
}
