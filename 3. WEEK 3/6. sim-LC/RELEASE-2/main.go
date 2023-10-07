package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		fmt.Println("RONALDO'S MINI SOCCER FIELD BOOKING ADMIN CLI")
		fmt.Println("Please select an option:")
		fmt.Println("1. Generate Revenue Report")
		fmt.Println("2. List Customers Without Payments")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1/2/3): ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			clearScreen()
			fmt.Println("Generating Revenue Report...\n")
			fmt.Println("| FIELD NAME       | TOTAL BOOKINGS | TOTAL REVENUE |")
			fmt.Println("| ---------------- | -------------- | ------------- |")
			fmt.Println("| Messi Park       | 125            | $8,750.00     |")
			fmt.Println("| Ronaldo Arena    | 100            | $5,000.00     |")
			fmt.Println("| Neymar Ground    | 80             | $3,600.00     |\n")
			fmt.Print("Press ENTER to return to the main menu...")
			reader.ReadString('\n')
		case "2":
			clearScreen()
			fmt.Println("Listing Customers Without Payments...\n")
			fmt.Println("| CUSTOMER NAME     | BOOKING ID | BOOKING DATE |")
			fmt.Println("| ----------------- | ---------- | ------------ |")
			fmt.Println("| Ahmad Setiawan    | 1023       | 2023-08-10   |")
			fmt.Println("| Siti Fatimah      | 1045       | 2023-08-12   |\n")
			fmt.Print("Press ENTER to return to the main menu...")
			reader.ReadString('\n')
		case "3":
			fmt.Println("Thank you for using Ronaldo's Admin CLI. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
