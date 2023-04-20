package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"Name"`
	School string `json:"School"`
	kelas  int
	umur   int
}

func main() {
	var jsonString = `{"Name": "Syaiful Imanudin", "school": "MAN 2 Kota Cirebon", "Kelas": 12, "Umur": 21}`
	var jsonData = []byte(jsonString)
	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("user :", data.Name)
	fmt.Println("school :", data.School)
	fmt.Println("kelas :", data.kelas)
	fmt.Println("umur :", data.umur)
}
