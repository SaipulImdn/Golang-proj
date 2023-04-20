package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Id    string
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

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var ID = "B001"
	err = db.
		QueryRow("select Name, Grade from tb_student where ID = ?", ID).
		Scan(&result.Name, &result.Grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.Name, result.Grade)

}

func main() {
	sqlQueryRow()
}
