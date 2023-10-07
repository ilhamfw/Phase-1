package main

import (
	"fmt"
)

func main() {
	// Data order dan employee statis
	orderData := [][]string{
		{"1", "1", "2823-88-07", "Served"},
		{"2", "1", "2023-08-07", "Preparing"},
		{"3", "3", "2023-08-08", "Preparing"},
		{"4", "4", "2023-08-08", "Ordered"},
		{"5", "5", "2023-08-08", "Preparing"},
	}

	employeeData := [][]string{
		{"John", "Doe", "Manager"},
		{"John", "Doe", "Manager"},
		{"Robert", "Brown", "Cook"},
		{"Jane", "Smith", "Waitress"},
		{"Steve", "Adam", "Manager"},
	}

	// Menampilkan data order
	fmt.Println("| OrderID | TableNumber | OrderDate  | Status     | FirstName | LastName | Position  |")
	fmt.Println("| ------- | ----------- | ---------- | ---------- | --------- | -------- | --------- |")

	for i, order := range orderData {
		// Pastikan indeks tidak melebihi jumlah data employee
		if i < len(employeeData) {
			fmt.Printf("| %-7s | %-11s | %-10s | %-10s | %-9s | %-8s | %-9s |\n", order[0], order[1], order[2], order[3], employeeData[i][0], employeeData[i][1], employeeData[i][2])
		} else {
			// Jika tidak ada data employee yang sesuai, tampilkan data kosong
			fmt.Printf("| %-7s | %-11s | %-10s | %-10s | %-9s | %-8s | %-9s |\n", order[0], order[1], order[2], order[3], "", "", "")
		}
	}

}
