package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlerwareAuth(handler)
	handler = MiddlerwareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler

	fmt.Println("server started at lcaolhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
}

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}

func MiddlerwareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlerwareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
