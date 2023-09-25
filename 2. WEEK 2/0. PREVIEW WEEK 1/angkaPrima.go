package main

import (
	"fmt"
	"math"
)

func main() {
	var angka int
	fmt.Print("Please enter a number: ")
	fmt.Scanln(&angka)

	if angka <= 0 {
		fmt.Println("Please enter a positive number greater than 0.")
		return
	}

	fmt.Println("the number is", checkEvenOdd(angka))

	if angka%2 == 0 {
		fmt.Println("Even numbers up to your input are:")
		for i := 2; i <= angka; i += 2 {
			fmt.Println(i)
		}
	} else {
		fmt.Println("Even numbers up to your input are:")
		for i := 2; i < angka; i += 2 {
			fmt.Println(i)
		}
	}

	if isPrime(angka) {
		fmt.Println("the number is prime number")
	} else {
		fmt.Println("the number is not a prime number")
	}
}

func checkEvenOdd(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
