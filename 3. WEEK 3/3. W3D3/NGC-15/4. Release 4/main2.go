package main

import (
	"fmt"
)

func main() {
	// Data order item statis
	orderItemData := [][]string{
		{"22", "1", "2024-08-07", "Preparing", "John", "Doe", "Manager", "Salad, Steak", "$35.00"},
		{"11", "1", "2024-88-87", "Served", "John", "Doe", "Manager", "Spaghetti Carbonara, Pasta", "$21.00"},
		{"33", "1", "2024-08-07", "Served", "Robert", "Brown", "Cook", "Salad", "$10.00"},
	}

	// Menampilkan data order item dalam format markdown
	fmt.Println("| OrderID | TableNumber | OrderDate  | Status     | FirstName | LastName | Position  | Items                             | TotalPrice |")
	fmt.Println("| ------- | ----------- | ---------- | ---------- | --------- | -------- | --------- | --------------------------------- | ---------- |")

	for _, orderItem := range orderItemData {
		fmt.Printf("| %-7s | %-11s | %-10s | %-10s | %-9s | %-8s | %-9s | %-33s | %-10s |\n", orderItem[0], orderItem[1], orderItem[2], orderItem[3], orderItem[4], orderItem[5], orderItem[6], orderItem[7], orderItem[8])
	}
}
