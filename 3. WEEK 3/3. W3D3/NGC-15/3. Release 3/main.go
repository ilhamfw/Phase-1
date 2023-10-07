package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "context"
	"time"
)

func main() {
    // Buat koneksi ke database MySQL
    //db, err := sql.Open("mysql", "root:@tcp(localhost:3307/tasty17)")
	db, err := sql.Open("mysql", "root@tcp(localhost:3307)/tasty17")


    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()

    // Buat konteks dengan timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Query untuk mengambil data dari tabel MenuItems
    query := db.Ex "SELECT item_id, name, description, price, category FROM menuitems"
    _, err = db.ExecContext "INSERT INTO MenuItems (name, description, price, category) VALUES (?,?,?,?,?)","Spaghetti Carbonara","Traditional Italian dish with eggs","
    ('Spaghetti Carbonara', 'Traditional Italian dish with eggs, cheese, pancetta, and pepper.', 12.58, 'Main Course'),
    ('Caesar Salad', 'Fresh lettuce with Caesar dressing, croutons and parmesan.', 6.00, 'Starter'),
    ('Tiramisu', 'Classic Italian dessert with coffee-soaked sponge and mascarpone.', 5.50, 'Dessert');"
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return
    } 
    defer rows.Close()

    // Iterasi melalui hasil query dan mencetak data
    // fmt.Printf("| %-7s | %-20s | %-64s | %-5s | %-12s |\n", "item_id", "name", "description", "Price", "Category")
    // fmt.Println("|---------| -------------------- | ----------------------------------------------------------------- | ------| ----------- |")
    for rows.Next() {
        var itemID int
        var name, description, category string
        var price float64
        err := rows.Scan(&itemID, &name, &description, &price, &category)
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return
        }
        // fmt.Printf("| %-7d | %-20s | %-64s | %-5.2f | %-12s |\n", itemID, name, description, price, category)
    }
    if err := rows.Err(); err != nil {
        fmt.Println("Error iterating through rows:", err)
        return
    }
}
