package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// NG Challenge 1 : Variabel 1
	var myNum int32 = 50
	fmt.Println(myNum)

	var myNum2 float32 = 51
	fmt.Println(myNum2)

	var myNumStr string = "50"
	fmt.Println(myNumStr)

	//NG Challenge 1 : Variabel 2
	var x int32 = 5
	var y int32 = 10
	var z int32 = x + y
	fmt.Println(z)

	//NG Challenge 1 : CLI
	fmt.Print("Masukkan nama: ")
	scanner := bufio.NewScanner(os.Stdin)

	// Baca input dari pengguna
	scanner.Scan()
	nama := scanner.Text()

	// Cetak pesan salam
	fmt.Printf("Hello %s\n", nama)

	//NG Challenge 1 : Array & Slice 1
	// Membuat slice "people" dengan nama-nama awal
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}

	// Menampilkan panjang slice "people" awal
	fmt.Printf("Panjang slice people: %d %v\n", len(people), people)

	// Menambahkan "Hank" dan "Marie" ke dalam slice "people"
	people = append(people, "Hank", "Marie")

	// Menampilkan panjang slice "people" setelah penambahan "Hank" dan "Marie"
	fmt.Printf("Menambahkan Hank dan Marie ...\n")
	fmt.Printf("Panjang slice people: %d %v\n", len(people), people)

	// Membuat Array dengan kapasitas 3 data yang berisi map
	array := []map[string]string{
		{"name": "Hank", "gender" : "M"},
		{"name": "Heisenberg", "gender" : "M"},
		{"name": "Skyler", "gender" : "F"},
	}

	//Menampilkan array di terminal
	fmt.Println("Array awal")
	fmt.Println(array)

	// Modifikasi array untuk menambahkan "Mrs" atau "Mr" sesuai dengan gender
	for i := 0; i < len(array); i++ {
		if array[i]["gender"] == "F" {
			array[i]["name"] = "Mrs " + array[i]["name"]
		} else if array[i]["gender"] == "M" {
			array[i]["name"] = "Mr " + array[i]["name"]
		}
	}

	// Menampilkan array yang sudah dimodifikasi ke dalam terminal
	fmt.Println("Array Setelah Modifikasi:")
	fmt.Println(array)
}
