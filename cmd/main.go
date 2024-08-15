package main

import (
	"httpmongodb/internal/db"
	"httpmongodb/internal/handlers"
	"log"
	"net/http"
)

const (
	host = "localhost:8080"
)

func main() {
	db.Connect()
	log.Printf("Start server host: %v", host)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{id}", handlers.GetHandler)
	mux.HandleFunc("GET /", handlers.GetAllHandler)
	mux.HandleFunc("POST /", handlers.PostHandler)
	mux.HandleFunc("PUT /{id}", handlers.PutHandler)
	mux.HandleFunc("DELETE /{id}", handlers.DeleteHandler)

	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatal(err)
	}
}
