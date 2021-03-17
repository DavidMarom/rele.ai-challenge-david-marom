package main

import (
	"log"
	"net/http"
	"yaml/yaml"
)

// Routes
func handleRequests() {
	http.HandleFunc("/ungroup", yaml.Ungroup)
	http.HandleFunc("/group", yaml.Group)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
