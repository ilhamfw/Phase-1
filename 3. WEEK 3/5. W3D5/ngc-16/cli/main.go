package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
    "ngc-16/config"
    "ngc-16/handler"
)

func main() {
    config.ConnectDB()

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
            handler.RegisterUser(config.DB)
        case 2:
            handler.LoginUser(config.DB)
        case 3:
            fmt.Println("Terima kasih! Keluar dari aplikasi.")
            os.Exit(0)
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}
