package main

import (
	"fmt"
)

// Tambah Struct
type Result struct {
	Operation string
	Value     float64
}

// 1. Buat 2 go routine untuk + dan -
func penjumlahan(chInput chan [2]float64, chOutput chan Result) {
	data := <-chInput
	chOutput <- Result{"tambah", data[0] + data[1]}
}

func pengurangan(chInput chan [2]float64, chOutput chan Result) {
	data := <-chInput
	chOutput <- Result{"kurang", data[0] - data[1]}
}

func perkalian(chInput chan [2]float64, chOutput chan Result) {
	data := <-chInput
	chOutput <- Result{"kali", data[0] * data[1]}
}

func pembagian(chInput chan [2]float64, chOutput chan Result) {
	data := <-chInput
	if data[1] != 0 {
		chOutput <- Result{"bagi", data[0] / data[1]}
	} else {
		chOutput <- Result{"bagi", 0.0} // Kita buat kembalian
	}
}

func main() {
	// 2. Channel lain
	chAddInput := make(chan [2]float64)
	chSubInput := make(chan [2]float64)
	chKalInput := make(chan [2]float64)
	chBagInput := make(chan [2]float64)
	chOutput := make(chan Result) // Perbaiki tipe channel

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
	for i := 0; i < 4; i++ {
		result := <-chOutput
		fmt.Printf("Hasil %s: %.2f\n", result.Operation, result.Value)
	}
}
