package main

import (
	"fmt"
	
)

// 1. Buat 2 go routine untuk + dan -
func penjumlahanBuffered(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	result := data[0] + data[1]
	chOutput <- result
}

func penguranganBuffered(chInput chan [2]float64, chOutput chan float64) {
	data := <-chInput
	result := data[0] - data[1] // Perbaiki tanda -
	chOutput <- result
}

func main() {
	// 2. Channel lain
	chAddInput := make(chan [2]float64, 1)
	chSubInput := make(chan [2]float64, 1)
	chOutput := make(chan float64, 2)

	// // 3. Inisialisasi go routine unbuffered
	// go penjumlahanUnBuffered(chAddInput, chOutput)
	// go penguranganUnBuffered(chSubInput, chOutput)

	// 3. Inisialisasi go routine Buffered
	go penjumlahanBuffered(chAddInput, chOutput)
	go penguranganBuffered(chSubInput, chOutput)

	// time.Sleep(100 * time.Millisecond) // Tunggu sejenjak untuk goroutine mulai lagi

	// 4. Kirim data ke go routine
	chAddInput <- [2]float64{5, 3}
	chSubInput <- [2]float64{5, 3} // Perbaiki tipe data channel

	// 5. Menerima dan mencetak hasil
	fmt.Println("Hasil Penjumlahan (Buffered):", <-chOutput)
	fmt.Println("Hasil Pengurangan (Buffered):", <-chOutput)
}
