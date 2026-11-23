package server

import (
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()
	port := "8080"

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /ascii-art", asciiHandler)

	// initialize the server here
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Println("Server starting on http://localhost:8080/home")
	log.Fatal(server.ListenAndServe())
}
