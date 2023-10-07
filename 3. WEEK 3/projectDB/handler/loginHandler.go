// Di handler/loginHandler.go
package handler

import (
    "projectDB/config"
    "projectDB/entity" // Mengganti entity dengan user sesuai dengan nama paket entitas Anda.
    "golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (user.User, bool) {
    var u user.User

    // Query ke database
    row := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=?", username)

    // Pindai hasil query ke struct user
    err := row.Scan(&u.ID, &u.Username, &u.Password)
    if err != nil {
        return u, false
    }

    // Bandingkan password yang di-hash dengan yang dimasukkan oleh pengguna
    err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    if err != nil {
        return u, false
    }

    return u, true
}
