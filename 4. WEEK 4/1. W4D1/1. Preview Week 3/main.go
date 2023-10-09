package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "root"
	dbHost = "localhost:3307"
	dbName = "db_bookstore"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(%s)/%s", dbUser, dbHost, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check for command input
	if len(os.Args) < 2 {
		fmt.Println("Please specify a command: books, sales, customers, topauthor")
		return
	}
	command := os.Args[1]

	switch command {
	case "books":
		listBooksByAuthor(db, "Jane Smith")
	case "sales":
		totalSalesByBookType(db)
	case "customers":
		identifyCustomersWithMultipleOrders(db)
	case "topauthor":
		displayTopEarningAuthor(db)
	default:
		fmt.Println("Unknown command:", command)
	}
}

func listBooksByAuthor(db *sql.DB, authorName string) {
	rows, err := db.Query("SELECT book_title FROM books WHERE author_id = (SELECT author_id FROM authors WHERE author_name = ?)", authorName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Books by", authorName, ":")
	for rows.Next() {
		var bookTitle string
		err := rows.Scan(&bookTitle)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bookTitle)
	}
}


func totalSalesByBookType(db *sql.DB) {
	rows, err := db.Query("SELECT b.book_type, SUM(o.price) as total_sales FROM orders o JOIN books b ON o.book_id = b.book_id GROUP BY b.book_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Total Sales by Book Type:")
	for rows.Next() {
		var bookType string
		var totalSales float64
		err := rows.Scan(&bookType, &totalSales)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: $%.2f\n", bookType, totalSales)
	}
}


func identifyCustomersWithMultipleOrders(db *sql.DB) {
	rows, err := db.Query("SELECT customer_name, COUNT(*) as order_count FROM orders JOIN customers ON orders.customer_id = customers.customer_id GROUP BY customer_name HAVING order_count > 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Customers with Multiple Orders:")
	for rows.Next() {
		var customerName string
		var orderCount int
		err := rows.Scan(&customerName, &orderCount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %d orders\n", customerName, orderCount)
	}
}

func displayTopEarningAuthor(db *sql.DB) {
	rows, err := db.Query("SELECT a.author_name, SUM(o.price) as earnings FROM authors a JOIN books b ON a.author_id = b.author_id JOIN orders o ON b.book_id = o.book_id GROUP BY a.author_name ORDER BY earnings DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Author with the Highest Earnings:")
	for rows.Next() {
		var authorName string
		var earnings float64
		err := rows.Scan(&authorName, &earnings)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: $%.2f\n", authorName, earnings)
	}
}

