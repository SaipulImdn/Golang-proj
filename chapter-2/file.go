package main

import (
	"fmt"
	"os"
)

var path = "Users/Public/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("=== file berhasil dibuat", path)
}

func main() {
	createFile()
}
