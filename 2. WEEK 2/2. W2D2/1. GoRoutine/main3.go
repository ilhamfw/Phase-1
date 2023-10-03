package main

import (
	"fmt"
	"time"
)

// 1. Definisikan Struct
type Notification struct {
	LogMessage  string
	Message string
}

// 2. Fungsi Asynchronous
func sendMessAsync(logMessage, message string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("%s: %s\n", logMessage, message)
}

func main() {
	// 3. Define dataset
	notifications := []Notification{
		{LogMessage: "[INFO]", Message: "User 123 logged in"},
		{LogMessage: "[WARN]", Message: "Memory usage is high"},
		{LogMessage: "[ERROR]", Message: "Failed to fetch data from API"},
		
	}

	// 4. Eksekusi Async
	for _, notification := range notifications {
		go sendMessAsync(notification.LogMessage, notification.Message)
	}

	// Tunggu beberapa saat agar goroutine berjalan
	fmt.Println("Application continues after logging...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main application finished.")
}

// Coba
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("Application continues after logging...")

// 	go func() {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println("[WARN] - Memory usage is high")
// 	}()

// 	go func() {
// 		time.Sleep(200 * time.Millisecond)
// 		fmt.Println("[INFO] - User 123 logged in")
// 	}()

// 	go func() {
// 		time.Sleep(300 * time.Millisecond)
// 		fmt.Println("[ERROR] - Failed to fetch data from API")
// 	}()

// 	// Tunggu sejenak sebelum pesan tercetak
// 	time.Sleep(400 * time.Millisecond)

// 	fmt.Println("Main application finished")
// }
