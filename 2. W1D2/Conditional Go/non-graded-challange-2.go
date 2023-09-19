package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func conditional1() {
	// Seed random number generator with current time
	rand.Seed(time.Now().UnixNano())

	// Membuat scanner untuk membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	// Meminta input nama dari pengguna
	fmt.Print("Masukkan Nama: ")
	scanner.Scan()
	nama := scanner.Text()

	// Meminta input angka dari pengguna
	fmt.Print("Masukkan Angka (1-100): ")
	scanner.Scan()
	angkaStr := scanner.Text()

	// Mengonversi input angka ke tipe int
	angka, err := strconv.Atoi(angkaStr)
	if err != nil {
		fmt.Println("Masukan angka yang valid.")
		return
	}

	// Menentukan pesan berdasarkan angka yang dihasilkan menggunakan switch case
	var pesan string
	switch {
	case angka > 80:
		pesan = "anda sangat beruntung"
	case angka > 60:
		pesan = "anda beruntung"
	case angka > 40:
		pesan = "anda kurang beruntung"
	default:
		pesan = "anda sial"
	}

	// Menampilkan pesan berdasarkan angka
	fmt.Printf("Selamat %s, %s\n", nama, pesan)
}
