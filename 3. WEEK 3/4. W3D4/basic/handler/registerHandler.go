package handler

import (
	"database/sql"
	"basic/entity"
)



// func RegisterStudent(db *sql.DB, student entity.Student) error {
// 	//_, err := db.Exec("INSERT INTO students (name, email) VALUES (?, ?)", student.Name, student.Email)
// 	_, err := db.Exec("INSERT INTO students (name, email) VALUES (?, ?)", student.Name, student.Email)

// 	return err
// }
func RegisterStudent(db *sql.DB, student entity.Student) error {
    // Gunakan prepared statement untuk menghindari SQL injection
    stmt, err := db.Prepare("INSERT INTO students (name, email) VALUES (?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(student.Name, student.Email)
    return err
}
