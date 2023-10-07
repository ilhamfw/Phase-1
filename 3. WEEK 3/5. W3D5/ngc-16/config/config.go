package config

import (
    "database/sql"
    "log"
    "ngc-16/config"
    "ngc-16/handler"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    DB, err = sql.Open("mysql", "root@tcp(localhost:3307)/w3d5")
    if err != nil {
        log.Fatal(err)
    }
}
