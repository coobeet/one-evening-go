package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", Hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, "+name)
	// Your solution goes here. Good luck!
}
