package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conditional2() {
	// Membuat scanner untuk membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	// Meminta input nama dari pengguna
	fmt.Print("Masukkan Nama: ")
	scanner.Scan()
	nama := scanner.Text()
	nama = strings.TrimSpace(nama)

	// Meminta input umur dari pengguna
	fmt.Print("Masukkan Umur: ")
	scanner.Scan()
	umurStr := scanner.Text()

	// Mengonversi input umur ke tipe data integer
	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("PESAN ERROR: Umur harus berupa angka.")
		return
	}

	// Memeriksa umur dan menampilkan pesan sesuai dengan ketentuan
	if umur < 0 || umur > 100 {
		fmt.Println("PESAN ERROR: Umur harus berada dalam rentang 0-100.")
	} else if umur > 18 {
		fmt.Println("Silahkan masuk, ", nama)
	} else {
		fmt.Println("Dilarang masuk, maksimal umur 19.")
	}
}

//explore strings string.TrimSpace
