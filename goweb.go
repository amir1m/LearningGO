package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
