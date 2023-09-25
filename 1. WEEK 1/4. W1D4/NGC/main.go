// package main

// import (
// 	"fmt"
// )

// type Person struct {
// 	Name string
// 	Age  int
// 	Job  string
// }

// // Method GetInfo untuk menampilkan informasi Person
// func (p *Person) GetInfo() {
// 	fmt.Printf("Name: %s\n", p.Name)
// 	fmt.Printf("Age: %d\n", p.Age)
// 	fmt.Printf("Job: %s\n", p.Job)
// }

// // Method AddYear untuk menambahkan usia dan mengubah pekerjaan jika usia >= 50
// func (p *Person) AddYear() {
// 	p.Age++
// 	if p.Age >= 50 {
// 		p.Job = "Retired"
// 	}
// }

// func main() {
// 	// Membuat instance Person dengan Age awal 45
// 	person := &Person{
// 		Name: "Bambang",
// 		Age:  45,
// 		Job:  "Progammer",
// 	}

// 	// Menampilkan informasi awal
// 	fmt.Println("Informasi Awal:")
// 	person.GetInfo()

// 	// Menambah usia 5 tahun
// 	for i := 0; i < 5; i++ {
// 		person.AddYear()
// 	}

// 	// Menampilkan informasi setelah penambahan usia
// 	fmt.Println("\nInformasi Setelah Penambahan Usia:")
// 	person.GetInfo()
// }

package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  string
}

// Method GetInfo untuk menampilkan informasi Person
func (p *Person) GetInfo() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Job: %s\n", p.Job)
}

func main() {
	// Membuat slice dari Person
	people := []Person{
		{Name: "Bambang", Age: 20, Job: "Progammer"},
		{Name: "Joko", Age: 30, Job: "Engineer"},
		{Name: "Siti", Age: 25, Job: "Teacher"},
	}

	// Melakukan pengulangan dan menampilkan informasi setiap Person
	for _, person := range people {
		person.GetInfo()
		fmt.Println()
	}
}
