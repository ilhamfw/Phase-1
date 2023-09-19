package main

import (
	"fmt"
	"sort"
	"strings"
)

func loop1() {
	// Membuat slice yang berisi data map orang-orang
	people := []map[string]interface{}{
		{"name": "Hank", "Age": 50, "Job": "Polisi"},
		{"name": "Heisenberg", "Age": 52, "Job": "Ilmuwan"},
		{"name": "Skyler", "Age": 48, "Job": "Akuntan"},
	}

	// Loop melalui slice dan tampilkan informasi setiap orang
	for _, person := range people {
		name, _ := person["name"].(string)
		age, _ := person["Age"].(int)
		job, _ := person["Job"].(string)

		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %d, dan saya bekerja sebagai %s\n", name, age, job)
	}
}

func loop2() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	// Fungsi untuk menghitung rata-rata dari slice
	average := func(numbers []float64) float64 {
		total := 0.0
		for _, num := range numbers {
			total += num
		}
		return total / float64(len(numbers))
	}

	// Fungsi untuk menghitung median dari slice
	median := func(numbers []float64) float64 {
		sort.Float64s(numbers)
		mid := len(numbers) / 2
		if len(numbers)%2 == 0 {
			return (numbers[mid-1] + numbers[mid]) / 2.0
		}
		return numbers[mid]
	}

	// Menghitung rata-rata, jumlah, dan median untuk slice1
	avg1 := average(slice1)
	sum1 := 0.0
	for _, num := range slice1 {
		sum1 += num
	}
	med1 := median(slice1)

	// Menghitung rata-rata, jumlah, dan median untuk slice2
	avg2 := average(slice2)
	sum2 := 0.0
	for _, num := range slice2 {
		sum2 += num
	}
	med2 := median(slice2)

	// Menampilkan hasil
	fmt.Printf("Slice 1: Rata-rata=%.2f, Jumlah=%.2f, Median=%.2f\n", avg1, sum1, med1)
	fmt.Printf("Slice 2: Rata-rata=%.2f, Jumlah=%.2f, Median=%.2f\n", avg2, sum2, med2)
}

func loop3() {
	// Meminta input dari pengguna
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	// Menghapus spasi dan mengubah huruf menjadi huruf kecil
	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))

	// Mengecek apakah kata merupakan palindrome
	result := isPalindrome(kata)

	// Menampilkan hasil
	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

// Fungsi untuk memeriksa apakah kata merupakan palindrome
func isPalindrome(kata string) bool {
	length := len(kata)
	for i := 0; i < length/2; i++ {
		if kata[i] != kata[length-i-1] {
			return false
		}
	}
	return true
}

func loop4() {
	// Meminta input dari pengguna
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	// Menghitung jumlah karakter 'x' dan 'o' dalam string
	countX := strings.Count(kata, "x")
	countO := strings.Count(kata, "o")

	// Menampilkan hasil
	result := countX == countO
	fmt.Println(result)
}

func loop5() {
	// Slice integer yang akan diurutkan
	slice := []int{5, 3, 8, 2, 1, 7, 4}

	// Proses pengurutan menggunakan Bubble Sort
	length := len(slice)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			// Jika elemen saat ini lebih kecil dari elemen berikutnya, tukar mereka
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	// Menampilkan slice yang sudah diurutkan dari besar ke kecil
	fmt.Println(slice)
}

func loop6() {
	// Input jumlah baris yang diinginkan
	var rows1 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows1)

	// Loop untuk mencetak barisan asterisks ke bawah
	for i := 0; i < rows1; i++ {
		fmt.Println("*")
	}
}

func loop7(){
	// Input jumlah baris yang diinginkan
	// Input jumlah baris yang diinginkan
	var rows2 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows2)

	// Loop untuk mencetak barisan asterisks dengan nested loop
	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("*")
		}
		fmt.Println() // Pindah ke baris baru setelah mencetak satu baris bintang
	}
}

func loop8(){
	// Input jumlah baris yang diinginkan
	var rows3 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows3)

	// Loop pertama untuk baris
	for i := 1; i <= rows3; i++ {
		// Loop kedua untuk mencetak simbol '*' sebanyak i kali
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println() // Pindah ke baris baru setelah mencetak satu baris bintang
	}
}

func loop9(){
	// Input jumlah baris yang diinginkan
	var rows4 int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&rows4)

	// Loop untuk mencetak barisan bintang dalam bentuk tangga terbalik
	for i := 0; i < rows4; i++ {
		// Loop dalam untuk mencetak bintang pada baris saat ini
		for j := 0; j < rows4-i; j++ {
			fmt.Print("*")
		}
		fmt.Println() // Pindah ke baris baru sebelum mencetak baris berikutnya
	}
}
