package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		body, err := ioutil.ReadAll(r.Body)
		_ = err
		_ = body
		time.Sleep(10 * time.Second)

		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request cancelled")
			} else {
				log.Println("unknow error occured.", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8000", nil)
}
