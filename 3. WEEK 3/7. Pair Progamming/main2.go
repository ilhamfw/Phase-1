package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Game struct {
	GameID      int
	Title       string
	Description string
	DatePublish string
	Rating      int
}

func clearScreen() {
	cmd := exec.Command("clear") // "clear" adalah perintah untuk membersihkan layar pada sistem macOS/Linux
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	defer func() {
		clearScreen()
		fmt.Println("Terminal has been cleared.")
	}()

	// Mengganti nilai sesuai dengan informasi MySQL Anda
	dsn := "root@tcp(localhost:3307)/game"

	// Membuka koneksi ke database MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Menjalankan program menu
	for {
		action, err := showMenu()
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "list-games":
			listGames(db)
		case "create-game":
			createGame(db)
		case "update-game":
			updateGame(db)
		case "delete-game":
			deleteGame(db)
		case "exit":
			os.Exit(0)
		}
	}
}

func showMenu() (string, error) {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"list-games", "create-game", "update-game", "delete-game", "exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func listGames(db *sql.DB) {
	fmt.Println("------------[LIST-GAMES]---------------------------")
	fmt.Printf("| %-6s | %-12s | %-20s | %-12s | %-6s |\n", "gameID", "title", "description", "date_publish", "rating")
	fmt.Println("| ------ | ------------ | ---------------------- | ------------ | ------ |")

	// Query database untuk membaca data game dan menampilkannya
	rows, err := db.Query("SELECT * FROM Game")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var game Game
		err := rows.Scan(&game.GameID, &game.Title, &game.Description, &game.DatePublish, &game.Rating)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("| %-6d | %-12s | %-20s | %-12s | %-6d |\n", game.GameID, game.Title, game.Description, game.DatePublish, game.Rating)
	}
}

func createGame(db *sql.DB) {
	fmt.Println("------------[CREATE-GAME]-------------------------")

	// Minta input dari pengguna untuk data game baru
	title := promptInput("Title: ")
	description := promptInput("Description: ")
	datePublish := promptInput("Date Publish (YYYY-MM-DD): ")
	rating := promptInput("Rating: ")

	// Query untuk menyisipkan game baru ke dalam database
	_, err := db.Exec("INSERT INTO Game (title, description, date_publish, rating) VALUES (?, ?, ?, ?)", title, description, datePublish, rating)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game berhasil dibuat!")
}


func updateGame(db *sql.DB) {
	fmt.Println("------------[UPDATE-GAME]-------------------------")

	// Minta input ID game yang ingin diperbarui
	gameID := promptInput("Game ID to update: ")

	// Minta input dari pengguna untuk data yang diperbarui
	title := promptInput("New Title: ")
	description := promptInput("New Description: ")
	datePublish := promptInput("New Date Publish (YYYY-MM-DD): ")
	rating := promptInput("New Rating: ")

	// Query untuk memperbarui game berdasarkan ID
	_, err := db.Exec("UPDATE Game SET title=?, description=?, date_publish=?, rating=? WHERE gameID=?", title, description, datePublish, rating, gameID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game berhasil diperbarui!")
}


func deleteGame(db *sql.DB) {
	fmt.Println("------------[DELETE-GAME]-------------------------")

	// Minta input ID game yang ingin dihapus
	gameID := promptInput("Game ID to delete: ")

	// Query untuk menghapus game berdasarkan ID
	_, err := db.Exec("DELETE FROM Game WHERE gameID=?", gameID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Game berhasil dihapus!")
}

func promptInput(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

