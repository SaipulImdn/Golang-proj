package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://www.google.com")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT ")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}

		w.Write([]byte("hello"))
	})

	log.Println("starting app at :9000")
	http.ListenAndServe(":9000", nil)
}
