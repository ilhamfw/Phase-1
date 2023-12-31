package main

import (
	"projectDB/config"
	"projectDB/entity"
	"projectDB/handler"
	"fmt"
)

func main() {
	// Menghubungkan ke database
	config.ConnectDB()

	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var username, password string
			fmt.Println("Registrasi:")
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)

			user := user.User{
				Username: username,
				Password: password, // Perbaiki penulisan 'password' menjadi 'Password'
			}

			err := handler.Register(user)
			if err != nil {
				fmt.Println("Kesalahan saat registrasi:", err)
			} else {
				fmt.Println("Registrasi berhasil")
			}

		case 2:
			var username, password string
			fmt.Println("Login:")
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)

			// Memanggil handler untuk login
			user, isAuthenticated := handler.Login(username, password)
			if isAuthenticated {
				fmt.Println("Login berhasil, selamat datang", user.Username)
			} else {
				fmt.Println("Login gagal")
			}

		case 3:
			fmt.Println("Keluar dari program")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
