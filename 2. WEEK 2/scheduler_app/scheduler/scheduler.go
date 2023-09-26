package scheduler

import (
	"fmt"
	"sync"
)

// ScheduleTask menjadwalkan dan menjalankan tugas secara asinkron dengan menggunakan goroutine.
func ScheduleTask(task func(*sync.WaitGroup)) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		task(&wg)
	}()
	wg.Wait()
}

// FibonacciTask adalah contoh fungsi Fibonacci yang akan dijadwalkan.
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

// AkronimTask adalah contoh fungsi Akronim yang akan dijadwalkan.
func AkronimTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Running Akronim Task")
	// Panggil fungsi Akronim di sini
}

// ScrabbleScoreTask adalah contoh fungsi Scrabble Score yang akan dijadwalkan.
func ScrabbleScoreTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Running Scrabble Score Task")
	// Panggil fungsi Scrabble Score di sini
}
