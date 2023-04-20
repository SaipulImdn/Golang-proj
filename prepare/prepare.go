package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	ID    string
	Name  string
	Age   int
	Grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:saipul1452@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select Name, Grade from tb_student where ID = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.Name, &result1.Grade)
	fmt.Printf("Name: %s\nGrade %d\n", result1.Name, result1.Grade)

	var result2 = student{}
	stmt.QueryRow("W001").Scan(&result2.Name, &result2.Grade)
	fmt.Printf("Name: %s\nGrade: %d\n", result1.Name, result2.Grade)

	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result3.Name, result3.Grade)
	fmt.Printf("Name: %s\nGrade: %d\n", result3.Name, result3.Grade)
}

func main() {
	sqlPrepare()
}
