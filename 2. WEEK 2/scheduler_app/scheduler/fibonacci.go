// package main

// import (
//     "fmt"
// )

// func fibonacci(n int) int {
//     if n <= 0 {
//         return 0
//     } else if n == 1 {
//         return 1
//     } else {
//         return fibonacci(n-1) + fibonacci(n-2)
//     }
// }

// func main() {
//     for i := 1; i <= 10; i++ {
//         result := fibonacci(i)
//         fmt.Printf("Bilangan Fibonacci ke-%d adalah %d\n", i, result)
//     }
// }

//Dengan Konkurensi
package scheduler

import (
	"fmt"
	"sync"
)

// FibonacciTask adalah fungsi untuk Fibonacci yang akan dijadwalkan.
func FibonacciTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("Enter the number: ")
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		result := fibonacci(i)
		fmt.Printf("Fibonacci number at position %d is %d\n", i, result)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

