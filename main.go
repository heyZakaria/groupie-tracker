package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

const port = "8080"

func main() {
	fmt.Println("http://localhost:" + port + "/")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/Artist", handlers.ArtistDetails)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
