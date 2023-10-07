package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:@tcp(localhost:3307)/tasty17"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// SQL statements for inserting data
	sqlStatements := []string{
		// Inserting a few employees
		`INSERT INTO Employees (first_name, last_name, position) VALUES 
			('John', 'Doe', 'Manager'),
			('Jane', 'Smith', 'Waitress'),
			('Robert', 'Brown', 'Cook');`,
		// Inserting a few menu items
		`INSERT INTO MenuItems (name, description, price, category) VALUES 
			('Spaghetti Carbonara', 'Traditional Italian dish with eggs, cheese, pancetta, and pepper.', 12.58, 'Main Course'),
			('Caesar Salad', 'Fresh lettuce with Caesar dressing, croutons and parmesan.', 6.00, 'Starter'),
			('Tiramisu', 'Classic Italian dessert with coffee-soaked sponge and mascarpone.', 5.50, 'Dessert');`,
		// Inserting a few orders
		`INSERT INTO Orders (table_number, employee_id, order_date, status) VALUES 
			(10, 1, '2023-08-09', 'Placed'),
			(5, 2, '2023-08-09', 'Completed');`,
		// Inserting items in the orders
		`INSERT INTO OrderItems (order_id, item_id, quantity, subtotal) VALUES 
			(1, 1, 2, 25.00), -- 2 Spaghetti Carbonara for the first order
			(2, 2, 1, 6.88), -- 1 Caesar Salad for the second order
			(2, 3, 1, 5.58); -- 1 Tiramisu for the second order;`,
		// Inserting payments for the orders
		`INSERT INTO Payments (order_id, payment_date, payment_method, total_amount) VALUES 
			(1, '2023-08-09', 'Credit Card', 25.00),
			(2, '2023-08-09', 'Cash', 11.50);`,
	}

	// Execute the SQL statements
	for _, sqlStatement := range sqlStatements {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Printf("Error executing SQL statement: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("All sample data created successfully!")
}
