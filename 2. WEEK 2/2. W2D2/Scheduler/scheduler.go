package scheduler

import (
	"fmt"
	"sync"
)

// ScheduleTask menjadwalkan dan menjalankan tugas secara asinkron dengan menggunakan goroutine.
func ScheduleTask(task func()) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		task()
	}()
	wg.Wait()
}

// FibonacciTask adalah contoh fungsi Fibonacci yang akan dijadwalkan.
func FibonacciTask() {
	fmt.Println("Running Fibonacci Task")
	// Panggil fungsi Fibonacci di sini
}

// AkronimTask adalah contoh fungsi Akronim yang akan dijadwalkan.
func AkronimTask() {
	fmt.Println("Running Akronim Task")
	// Panggil fungsi Akronim di sini
}

// ScrabbleScoreTask adalah contoh fungsi Scrabble Score yang akan dijadwalkan.
func ScrabbleScoreTask() {
	fmt.Println("Running Scrabble Score Task")
	// Panggil fungsi Scrabble Score di sini
}
