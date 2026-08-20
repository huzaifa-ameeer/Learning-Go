package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	_, _ = w.Write([]byte("Welcome, try to /login?name=Huzaifa"))

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	_, _ = w.Write([]byte(name))
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server running")
	http.ListenAndServe(":3000", nil)

}
