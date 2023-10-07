// handler/loginHandler.go

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

func LoginUser(db *sql.DB) {
    fmt.Println("== Login User ==")
    fmt.Print("Email: ")
    var email, password string
    fmt.Scan(&email)
    fmt.Print("Password: ")
    fmt.Scan(&password)

    // Ambil data pengguna dari database berdasarkan email
    var storedPassword string
    err := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
    if err != nil {
        log.Fatal("Gagal mengambil data pengguna:", err)
    }

    // Bandingkan password yang di-hash dengan yang dimasukkan oleh pengguna
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
    if err != nil {
        fmt.Println("Login gagal. Email/Password salah.")
        return
    }

    fmt.Println("Sukses Login!")
}
