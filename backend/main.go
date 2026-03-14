package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	// API endpoints
	http.HandleFunc("/builds", GetAllBuildsHandler)
	http.HandleFunc("/build/create", CreateBuildHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}