package main

import (
	"database/sql"
	"fmt"
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
)
func main(){
	// Buat koneksi ke database mysql
	// db, err := sql.Open("mysql","username:password@tcp(localhost:3307/dbname)")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/sesi7")


	//handling error
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// buat test koneksi
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect the database")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Succesfully connect to the database")

	// 1. Buat Operasi Table dan Insert sample data dengan fungsi EXEC()

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS test(id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50))")

	// 2. Cek Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("Succesfully created Table")
	// }
	// 3. Insert Sample Data
	// _, err = db.Exec("INSERT INTO test (name) VALUES (?)","John")

	// // 4. Cek Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("Succesfully Insert Data")
	// }

	// 5. Cara penggunaan Fungsi Query() untuk menampilkan data (SELECT)

	// rows, err := db.Query("SELECT * FROM test")

	// // 6. Cek Error
	// if err != nil(
	// 	fmt.Println(err.Error())
	// 	return
	// )
	// //7. tutup koneksi
	// defer rows.Close()

	// // 8. Perlu lakukan iterasi
	// for rows.Next(){
	// 	var id int
	// 	var name string
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	} 
	// 	fmt.Println(id, name)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5 +time.Second)
	defer cancel()

	// create table
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS user (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50))")

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Succesfully created Table")
	}

	// INsert Data
	_, err = db.ExecContext(ctx, "INSERT INTO user (name) VALUES (?)", "Jane Smith")

	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Succesfully Insert Data")
	}

	// Select data
	rows, err := db.QueryContext(ctx, "SELECT * FROM user")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	fmt.Println("Listening All Record :")

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Println("ID:", id, "NAME:", name)
	}
	
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating through rows:", err)
		return
	}
	
}