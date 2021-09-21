package main

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to financier, by Nick Ingram."))
}

func main() {
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
