package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func group(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := reqBody

	w.Write([]byte(result))

}

func ungroup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := reqBody

	w.Write([]byte(result))
}

// Routes
func handleRequests() {
	http.HandleFunc("/ungroup", ungroup)
	http.HandleFunc("/group", group)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
