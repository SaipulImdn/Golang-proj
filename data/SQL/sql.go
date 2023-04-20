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

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", Age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)

	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
	}
}

func main() {
	sqlQuery()
}
