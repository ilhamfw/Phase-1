package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    Email    string
    FullName string
    Password []byte
    DOB      string
}

var users map[string]User

func main() {
    users = make(map[string]User)

    // Konfigurasi koneksi ke database MySQL
    db, err := sql.Open("mysql", "root@tcp(localhost:3307)/w3d5")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Mengecek koneksi ke database
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Terhubung ke database MySQL")

    // Jalankan program registrasi dan login di sini
    for {
        fmt.Println("Pilih aksi:")
        fmt.Println("1. Register User")
        fmt.Println("2. Login User")
        fmt.Println("3. Keluar")
        fmt.Print("Masukkan pilihan (1/2/3): ")

        var choice int
        _, err := fmt.Scanf("%d", &choice)
        if err != nil {
            fmt.Println("Input tidak valid. Silakan coba lagi.")
            continue
        }

        switch choice {
        case 1:
            registerUser(db)
        case 2:
            loginUser(db)
        case 3:
            fmt.Println("Terima kasih! Keluar dari aplikasi.")
            os.Exit(0)
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

func registerUser(db *sql.DB) {
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



func loginUser(db *sql.DB) {
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

func hashPassword(password string) ([]byte, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal("Gagal menghash password:", err)
    }
    return hashedPassword, nil
}
