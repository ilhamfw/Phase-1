package handler

import (
    "projectDB/config"
    "projectDB/entity" // Sesuaikan nama paket entitas dengan yang Anda gunakan
    "golang.org/x/crypto/bcrypt"
)

func Register(user user.User) error { // Gunakan user.User sesuai dengan nama entitas yang diimpor
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    _, err = config.DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)", user.Username, string(hashedPassword))
    return err
}
