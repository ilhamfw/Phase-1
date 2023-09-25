package main

import (
	"fmt"
	
)

func main() {

	// Int, string, Boolean
	// Pake Type Data
	// var name   = "Eko"
	// var age  = 19

	// Short Declaration
	name := "wahyu"
	age := 18

	// fmt.Println("Nama Sebelum = ", name)
	// name = "Abdul"

	// fmt.Println("Nama Sesudah = ", name)
	// fmt.Println("Umur = ", age)

	// Cek type data
	fmt.Printf("type data %T value %v\n", name, name)
	fmt.Printf("type data %T value %v\n", age, age)

	// Gak Pake Type data
	// var name  = "Eko"
	// var age = 19

	// fmt.Println("Nama Sebelum = ", name)
	// name = "Abdul"
	
	// fmt.Println. %
	// fmt.Println("Nama Sesudah = ", name)
	// fmt.Println("Umur = ", age)

	//Multiple Deklarasi
	var student1, student2, student3 string = "Fikri", "Daniel", "Dani"
	fmt.Println(student1, student2, student3, "<----Student Ftgo before")
	student1 = "Ilham"
	student2, student3 = "Nafisa", "osvaldo"
	fmt.Println(student1, student2, student3, "<----Student Ftgo after")

	name1, isStudent, age1 := "Risky", true, 18
	fmt.Println(name1, isStudent, age1)
	fmt.Printf("%T - %T - %T", name1, isStudent, age1)
	
	// Type data : Number (Integer)
	var salary1 = 30
	var salary2 = 30
	fmt.Println(salary1, salary2)
	fmt.Printf("%T - %T\n", salary1, salary2)

	// Boolean
	var isManager bool = true
	fmt.Printf("%T - %v", isManager, isManager)

	//String
	var lyric = `
	Menangislah
	Kan kau juga manusia
	Mana ada yang bisa
	Berlarut-larut
	Berpura-pura sempurna	
	`

	var query = `
	Select * FROM students
	Join courses On Student.id = course,student_id
	where student.id`
	fmt.Println(lyric, query)

	//Constan Operator
	const address string = "https://hacktiv8.com"

	fmt.Println(address)

	//Operator Aritmatika

	result := 40 + 30 - 20 * 2
	result = result + 1
	result += 1
	fmt.Println(result)

	mod := 5 % 2
	fmt.Println("Mod adalah : ",mod)

	//Operator Pembanding

	fmt.Println(result == mod, " = result > mod")

	//Operator Logika

	attitude := "AB"
	score := 75
	naikKelas := attitude == "A" && score > 70

	fmt.Println(naikKelas, "<--- Naik kelas")
	fmt.Println(!naikKelas, "<--- naik kelas negasi")


}