package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Pilih operasi aritmatika:")
	fmt.Println("> Penjumlahan (+) [a]")
	fmt.Println("> Pengurangan (-) [b]")
	fmt.Println("> Perkalian (*) [c]")
	fmt.Println("> Pembagian (/) [d]")

	var operation string
	fmt.Print("Pilihan: ")
	fmt.Scan(&operation)

	var result float64

	switch operation {
	case "a":
		result = doOperation(add)
	case "b":
		result = doOperation(subtract)
	case "c":
		result = doOperation(multiply)
	case "d":
		result = doOperation(divide)
	default:
		fmt.Println("Operasi tidak valid.")
		return
	}

	fmt.Printf("Hasil: %.2f\n", result)
}

func doOperation(op func(float64, float64) float64) float64 {
	var num1, num2 float64
	fmt.Print("Masukkan angka pertama: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Input tidak valid.")
		return 0
	}

	fmt.Print("Masukkan angka kedua: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Input tidak valid.")
		return 0
	}

	return op(num1, num2)
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		panic("Pembagian oleh nol tidak diizinkan")
	}
	return a / b
}
