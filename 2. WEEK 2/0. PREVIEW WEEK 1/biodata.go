package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var temanKelas = []Teman{
	{"Nama Teman 1", "Alamat Teman 1", "Pekerjaan Teman 1", "Alasan Teman 1"},
	{"Nama Teman 2", "Alamat Teman 2", "Pekerjaan Teman 2", "Alasan Teman 2"},
	// Tambahkan data teman-teman lainnya di sini
}


func biodata() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Penggunaan: go run biodata.go [nomor absen]")
		return
	}

	nomorAbsen := args[1]
	nomor, err := strconv.Atoi(nomorAbsen)
	if err != nil || nomor < 1 || nomor > len(temanKelas) {
		fmt.Println("Nomor absen tidak valid. Masukkan nomor absen yang benar.")
		return
	}

	teman := temanKelas[nomor-1]
	fmt.Printf("Data Teman %d:\n", nomor)
	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan: %s\n", teman.Alasan)
}
