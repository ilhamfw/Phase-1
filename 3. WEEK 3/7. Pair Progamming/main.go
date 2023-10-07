package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Mengganti nilai sesuai dengan informasi MySQL Anda
	dsn := "root@tcp(localhost:3307)/game"

	// Membuka koneksi ke database MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat tabel Game jika belum ada
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS Game (
		gameID INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		date_publish DATE,
		rating INTEGER
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Menjalankan program menu
	for {
		fmt.Println("\nGame Database Menu:")
		fmt.Println("1. Create Game")
		fmt.Println("2. Read Games")
		fmt.Println("3. Update Game")
		fmt.Println("4. Delete Game")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice (1/2/3/4/5): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case 1:
			createGame(db)
		case 2:
			readGames(db)
		case 3:
			updateGame(db)
		case 4:
			deleteGame(db)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func createGame(db *sql.DB) {
	fmt.Println("\nCreate Game:")
	fmt.Print("Enter title: ")
	var title string
	_, err := fmt.Scan(&title)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter description: ")
	var description string
	_, err = fmt.Scan(&description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter date_publish (YYYY-MM-DD): ")
	var datePublish string
	_, err = fmt.Scan(&datePublish)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter rating: ")
	var rating int
	_, err = fmt.Scan(&rating)
	if err != nil {
		log.Fatal(err)
	}

	insertGameSQL := `INSERT INTO Game (title, description, date_publish, rating) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(insertGameSQL, title, description, datePublish, rating)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game created successfully.")
}

func readGames(db *sql.DB) {
	fmt.Println("\nGames List:")
	rows, err := db.Query("SELECT * FROM Game")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("%-7s %-25s %-40s %-12s %-7s\n", "gameID", "title", "description", "date_publish", "rating")
	fmt.Println("---------------------------------------------------------")
	for rows.Next() {
		var gameID int
		var title, description, datePublish string
		var rating int
		if err := rows.Scan(&gameID, &title, &description, &datePublish, &rating); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-7d %-25s %-40s %-12s %-7d\n", gameID, title, description, datePublish, rating)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func updateGame(db *sql.DB) {
	fmt.Println("\nUpdate Game:")
	fmt.Print("Enter gameID to update: ")
	var gameID int
	_, err := fmt.Scan(&gameID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter new title: ")
	var title string
	_, err = fmt.Scan(&title)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter new description: ")
	var description string
	_, err = fmt.Scan(&description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter new date_publish (YYYY-MM-DD): ")
	var datePublish string
	_, err = fmt.Scan(&datePublish)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter new rating: ")
	var rating int
	_, err = fmt.Scan(&rating)
	if err != nil {
		log.Fatal(err)
	}

	updateGameSQL := `UPDATE Game SET title=?, description=?, date_publish=?, rating=? WHERE gameID=?`
	_, err = db.Exec(updateGameSQL, title, description, datePublish, rating, gameID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game updated successfully.")
}

func deleteGame(db *sql.DB) {
	fmt.Println("\nDelete Game:")
	fmt.Print("Enter gameID to delete: ")
	var gameID int
	_, err := fmt.Scan(&gameID)
	if err != nil {
		log.Fatal(err)
	}

	deleteGameSQL := `DELETE FROM Game WHERE gameID=?`
	_, err = db.Exec(deleteGameSQL, gameID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game deleted successfully.")
}
