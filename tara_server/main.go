package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" {
		name = "Guest"
	}
	if age == "" {
		age = "Unknown"
	}

	fmt.Fprintf(w, "Hi %s!\n", name)
	fmt.Fprintf(w, "Welcome to Tara's server!, you are %s years old.", age)
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
