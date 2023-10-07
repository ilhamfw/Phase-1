package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Gantilah dengan informasi koneksi database Anda
	dsn := "root:@tcp(localhost:3307)/tasty17"


	// Membuka koneksi ke database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Pernyataan DDL untuk membuat tabel-tabel
	ddlStatements := []string{
		`CREATE TABLE Employees (
			employee_id INT PRIMARY KEY AUTO_INCREMENT,
			first_name VARCHAR(50),
			last_name VARCHAR(50),
			position VARCHAR(50)
		);`,
		`CREATE TABLE MenuItems (
			item_id INT PRIMARY KEY AUTO_INCREMENT,
			description VARCHAR(50),
			price DECIMAL(10, 2),
			category VARCHAR(50)
		);`,
		`CREATE TABLE Orders (
			order_id INT PRIMARY KEY AUTO_INCREMENT,
			table_number INT,
			employee_id INT,
			order_date DATE, 
			status VARCHAR(50),
			FOREIGN KEY (employee_id) REFERENCES Employees(employee_id)
		);`,
		`CREATE TABLE OrderItems (
			order_item_id INT PRIMARY KEY AUTO_INCREMENT,
			order_id INT,
			item_id INT,
			quantity INT,
			subtotal DECIMAL(10, 2),
			FOREIGN KEY (order_id) REFERENCES Orders(order_id),
			FOREIGN KEY (item_id) REFERENCES MenuItems(item_id)
		);`,
		`CREATE TABLE Payments (
			payment_id INT PRIMARY KEY AUTO_INCREMENT,
			order_id INT,
			payment_date DATE,
			payment_method VARCHAR(50),
			total_amount DECIMAL(10, 2),
			FOREIGN KEY (order_id) REFERENCES Orders(order_id)
		);`,
	}

	// Loop melalui pernyataan DDL dan eksekusi masing-masing
	for _, ddl := range ddlStatements {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("Error executing DDL statement: %v\n", err)
		}
	}

	fmt.Println("All Tables created successfully")
}
