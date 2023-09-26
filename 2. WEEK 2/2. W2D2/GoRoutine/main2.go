package main

import (
	"fmt"
	"time"
)

// 1. Definisikan Struct
type Notification struct {
	UserID  int
	Message string
}

// 2. Fungsi Synchronous
func sendEmail(userID int, message string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Email Notification sent to user %d: %s\n", userID, message)
}

func main() {
	// 3. Define dataset
	notifications := []Notification{
		{UserID: 101, Message: "Your Order has been confirmed"},
		{UserID: 202, Message: "Your Order has been confirmed"},
		{UserID: 303, Message: "Your Order has been confirmed"},
	}

	// 4. Eksekusi Synchronous
	for _, notification := range notifications {
		sendEmail(notification.UserID, notification.Message)
	}

	// Tunggu beberapa saat agar goroutine berjalan
	fmt.Println("Main application continues.")
	time.Sleep(3 * time.Second)
	fmt.Println("Main application finished.")
}
