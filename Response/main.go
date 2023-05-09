package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("server at ready localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name   string
		Age    int
		Gender string
	}{
		{"Syaiful Imanudin", 21, "laki-laki"},
		{"Jason", 23, "laki-laki"},
		{"Tim Drake", 22, "laki-laki"},
		{"Damian Wayne", 21, "laki-laki"},
	}
	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(jsonInBytes)
}
