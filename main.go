package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting() string {
	return "Hello, World! 222"
}

func demo() string {
	return "demo"
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, greeting())
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, demo())
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/demo", demoHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
