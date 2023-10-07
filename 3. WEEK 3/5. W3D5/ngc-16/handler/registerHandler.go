// handler/registerHandler.go

package handler

import (
    "fmt"
    "log"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
    "ngc-16/config"
    "ngc-16/entity"
)

func RegisterUser(db *sql.DB) {
    fmt.Println("== Registrasi User ==")
    fmt.Print("Email: ")
    var email, fullName, dob, password string
    fmt.Scan(&email)
    fmt.Print("Full Name: ")
    fmt.Scan(&fullName)
    fmt.Print("Date of Birth: ")
    fmt.Scan(&dob)
    fmt.Print("Password: ")
    fmt.Scan(&password)

    // Hash password sebelum menyimpannya ke database
    hashedPassword, err := hashPassword(password)
    if err != nil {
        log.Fatal("Gagal menghash password:", err)
    }

    // Masukkan data pengguna ke dalam tabel "users" dalam database
    _, err = db.Exec("INSERT INTO users (email, full_name, dob, password) VALUES (?, ?, ?, ?)", email, fullName, dob, hashedPassword)
    if err != nil {
        log.Fatal("Gagal menyimpan data pengguna:", err)
    }

    fmt.Println("Registrasi berhasil.")
}

func hashPassword(password string) ([]byte, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal("Gagal menghash password:", err)
    }
    return hashedPassword, nil
}
