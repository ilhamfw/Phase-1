// package main

// import "fmt"

// // QUIZ 1

// func main() {
// 	// Deklarasi Variable Nama dan Umur
// 	nama := "Ilham"
// 	age := 20

// 	// Cetak Variable ke konsol
// 	fmt.Printf("Nama: %s\n", nama)
// 	fmt.Printf("Umur: %d\n", age)
// }

// QUIZ 2


// package main

// import "fmt"

// func main(){
// 	//  membuat map buah dan harga perkilo
// 	buah := make(map[string]float64)

// 	// Mengisi Map denga  data buah dan harga perkilo
// 	{
// 		buah["Apel"] = 10.5
// 		buah["Pisang"] = 8.5
// 		buah["Mangga"] = 15.5
// 	}

// 	// Menampilkan isi data buah dan harga perkilogram
// 	for nama, harga := range buah {
// 		fmt.Printf("Buah : %s, Harga perkilo : %.2f\n", nama,harga)
// 	}

// }

//QUIZ 3
// package main

// import "fmt"

// func main(){
// 	number := 8
// 	//Tulis disini kode untuk memeriksa apakah genap atau ganjil
// 	if number%2 == 0{
// 		fmt.Println("Ini Bilangan Genap")
// 	} else {
// 		fmt.Println("Ini bilangan Ganil")
// 	}
// }

//QUIZ 4
// package main

// import "fmt"

// func main(){
// 	var bulan int
// 	fmt.Print("Masukkan angka bulan (1-12): ")
// 	fmt.Scanln(&bulan)

// 	// Switch Case untuk mengembalikan nama bulan
// 	switch bulan {
// 	case 1:
// 		fmt.Println("Januari")
// 	case 2:
// 		fmt.Println("Februari")
// 	case 3:
// 		fmt.Println("Maret")
// 	case 4:
// 		fmt.Println("April")
// 	case 5:
// 		fmt.Println("Mei")
// 	case 6:
// 		fmt.Println("Juni")
// 	case 7:
// 		fmt.Println("Juli")
// 	case 8:
// 		fmt.Println("Agustus")
// 	case 9:
// 		fmt.Println("September")
// 	case 10:
// 		fmt.Println("Oktober")
// 	case 11:
// 		fmt.Println("November")
// 	case 12:
// 		fmt.Println("Desember")
// 	default:
// 		fmt.Println("Bulan tidak valid")
// 	}
// }

//QUIZ 5
// package main

// import "fmt"

// func main() {
//     for i := 1; i <= 10; i++ {
//         fmt.Println(i)
//     }
// }

//QUIZ 6
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var N int
// 	fmt.Print("Masukkan nilai N: ")
// 	fmt.Scanln(&N)

// 	fmt.Println("N suku pertama dari deret geometri dengan rasio 2:")
// 	suku := 1
// 	for i := 0; i < N; i++ {
// 		fmt.Printf("%d ", suku)
// 		suku *= 2
// 	}
// 	fmt.Println()
// }

//QUIZ 7
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Mencetak angka dari 1 hingga 20 dengan FizzBuzz:")
// 	for i := 1; i <= 20; i++ {
// 		switch {
// 		case i%3 == 0 && i%5 == 0: //Habis dibagi 3 dan 5 :FizzBuzz
// 			fmt.Println("FizzBuzz")
// 		case i%3 == 0: //Habis dibagi 3 : Fizz
// 			fmt.Println("Fizz")
// 		case i%5 == 0: //Habis dibagi 5 : Buzz
// 			fmt.Println("Buzz")
// 		default:
// 			fmt.Println(i)
// 		}
// 	}
// }

//QUIZ 8
// package main

// import "fmt"

// func main() {
// 	var N int
// 	fmt.Print("Masukkan nilai N: ")
// 	fmt.Scanln(&N)

// 	fmt.Println("Mencetak bilangan dari 1 hingga", N, "dengan FizzBuzz:")
// 	for i := 1; i <= N; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

//QUIZ 9
// package main

// import "fmt"

// func main() {
//     var N int
//     fmt.Print("Masukkan nilai N: ")
//     fmt.Scanln(&N)

//     fmt.Println("Mencetak pola kotak:")
//     for i := 1; i <= N; i++ {
//         for j := 1; j <= N; j++ {
//             if i == 1 || i == N || j == 1 || j == N {
//                 fmt.Print("*")
//             } else {
//                 fmt.Print(" ")
//             }
//         }
//         fmt.Println()
//     }
// }

//QUIZ 10
// package main

// import "fmt"

// func main() {
//     tinggi := 5

//     for i := 0; i < tinggi; i++ {
//         nilai := 1
//         for j := 0; j <= i; j++ {
//             fmt.Print(nilai, " ")
//             nilai = nilai * (i - j) / (j + 1)
//         }
//         fmt.Println()
//     }
// }

//QUIZ 11
// package main

// import "fmt"

// //Fungsi untuk menghitung jumlah bilangan
// func tambah(a, b int) int{
// 	return a + b
// }

// func main(){
// 	bilangan1 := 5
// 	bilangan2 := 7

// 	//memanggil fungsi tambah dan menampilkan hasilnya
// 	hasil := tambah(bilangan1, bilangan2)

// 	fmt.Printf("Hasil dari penjumlahan %d + %d = %d\n", bilangan1, bilangan2, hasil  )
// }

//QUIZ 12
// package main

// import "fmt"

// // Fungsi untuk menghitung jumlah dua bilangan bulat
// func tambah(a int, b int) int {
//     return a + b
// }

// func main() {
//     bilangan1 := 5
//     bilangan2 := 7

//     // Memanggil fungsi tambah dan menyimpan hasilnya dalam variabel hasil
//     hasil := tambah(bilangan1, bilangan2)

//     fmt.Printf("Hasil penjumlahan %d + %d = %d\n", bilangan1, bilangan2, hasil)
// }

//QUIZ 13

// package main

// import "fmt"

// // Fungsi VARIADIC untuk menghitung jumlah bilangan bulat
// func jumlahBilangan(bilangan ...int) int {
//     total := 0
//     for _, angka := range bilangan {
//         total += angka
//     }
//     return total
// }

// func main() {
//     hasil := jumlahBilangan(5, 7, 10, 15, 20)

//     fmt.Printf("Hasil penjumlahan adalah %d\n", hasil)
// }

//QUIZ 14
// package main

// import "fmt"

// // Fungsi variadic untuk mencari nilai maksimum
// func nilaiMaksimum(bilangan ...int) int {
//     if len(bilangan) == 0 {
//         // Mengembalikan 0 jika tidak ada argumen
//         return 0
//     }

//     maksimum := bilangan[0] // Inisialisasi dengan nilai pertama

//     for _, angka := range bilangan {
//         if angka > maksimum {
//             maksimum = angka
//         }
//     }
//     return maksimum
// }

// func main() {
//     hasil := nilaiMaksimum(5, 7, 10, 15, 20, 12, 8)

//     fmt.Printf("Nilai maksimum adalah %d\n", hasil)
// }

//QUIZ 15
// package main

// import (
//     "fmt"
// )

// type Book struct {
//     // Isikan kode funsi di sini
//     Title string
//     Author string
//     Price float64
// }

// func totalHarga(books []Book)float64{
//     // tulis kode fungsi di sini
//     total := 0.0
//     for _, book := range books {
//         total += book.Price
//     }
//     return total
// }

// func main() {
//     // Inisialisasi slice dari struct Book dan uji fungsi totalHarga
//     books := []Book{
//         {"Book 1", "Author 1", 12.99},
//         {"Book 2", "Author 2", 9.99},
//         {"Book 3", "Author 3", 15.99},
//     }

//     //uji fungsi totalHarga
//     total := totalHarga(books)
//     fmt.Printf("Total Harga Buku: $%.2f\n", total)
// }

//Quiz 16
package main

import "fmt"

func hitungBonus(perform string, gaji float64)float64{
    //tulis kode di sini
    switch perform {
    case "A":
        return 0.2 * gaji // Bonus 20%
    case "B":
        return 0.1 * gaji // Bonus 10%
    case "C":
        return 0.05 * gaji // Bonus 5%
    default:
        return 0.0 // Tidak ada bonus
    }
}

func main(){
    // Uji fungsi hitungBonus dengan beberapa kasus
    gaji := 1000.0
    perform := "A"
    bonus := hitungBonus(perform, gaji)
    fmt.Printf("Bonus untuk performa %s: $%.2f\n", perform, bonus)

    perform = "B"
    bonus = hitungBonus(perform, gaji)
    fmt.Printf("Bonus untuk performa %s: $%.2f\n", perform, bonus)

    perform = "C"
    bonus = hitungBonus(perform, gaji)
    fmt.Printf("Bonus untuk performa %s: $%.2f\n", perform, bonus)

    perform = "D"
    bonus = hitungBonus(perform, gaji)
    fmt.Printf("Bonus untuk performa %s: $%.2f\n", perform, bonus)

}





