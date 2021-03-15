package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func group(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	// w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	// result := "type: DataFile\npayload:\n████\n"
	result := ""

	docsArray := strings.Split(string(reqBody), "---")

	toLines := strings.Split(docsArray[0], "\n")
	for i = 2; i < len(toLines); i++ {
		result += toLines[i]
	}

	toLines = strings.Split(docsArray[1], "\n")
	for i = 3; i < len(toLines); i++ {
		result += toLines[i]
	}

	toLines = strings.Split(docsArray[2], "\n")
	for i = 3; i < len(toLines); i++ {
		result += toLines[i]
	}

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
	// Printf(aaa())
}
