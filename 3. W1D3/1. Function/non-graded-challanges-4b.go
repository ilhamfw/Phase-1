package main

import (
	"fmt"
)

// Fungsi untuk menghitung bilangan Fibonacci ke-n
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main2() {
	// Memanggil fungsi Fibonacci dengan nilai n
	n := 10 // Ganti dengan nilai n yang diinginkan
	result := Fibonacci(n)

	fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", n, result)
}
