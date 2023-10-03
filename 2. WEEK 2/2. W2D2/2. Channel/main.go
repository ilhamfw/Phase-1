//ch <- 5 // kita mengirim nilai 5 ke channel
// value := <-ch // menerima nilai dan kita simpan dalam variabel

//kalkulator
package main

import (
	"fmt"
)

// 1. Buat 2 go routine untuk + dan -
func penjumlahan(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	result := data[0] + data[1]
	chOutput <- result
}

func pengurangan(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	result := data[0] - data[1]
	chOutput <- result
}

func perkalian(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	result := data[0] * data[1]
	chOutput <- result
}

func pembagian(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	result := data[0] / data[1]
	chOutput <- result
}

func main() {
	// 2. Channel lain
	chAddInput := make(chan [2]float64)
	chSubInput := make(chan [2]float64)
	chKalInput := make(chan [2]float64)
	chBagInput := make(chan [2]float64)
	chOutput := make(chan float64)

	// 3. Inisialisasi go routine
	go penjumlahan(chAddInput, chOutput)
	go pengurangan(chSubInput, chOutput)
	go perkalian(chKalInput, chOutput) // Perbaiki nama channel
	go pembagian(chBagInput, chOutput) // Perbaiki nama channel

	// 4. Kirim data ke go routine
	chAddInput <- [2]float64{5, 3}
	chSubInput <- [2]float64{5, 3}
	chKalInput <- [2]float64{5, 3}
	chBagInput <- [2]float64{5, 3}

	// 5. Menerima dan mencetak hasil
	fmt.Println("Hasil Penjumlahan:", <-chOutput)
	fmt.Println("Hasil Pengurangan:", <-chOutput)
	fmt.Println("Hasil Perkalian:", <-chOutput)
	fmt.Println("Hasil Pembagian:", <-chOutput)
}
