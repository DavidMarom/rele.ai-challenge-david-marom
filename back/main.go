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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	j := 0
	result := "type: DataFile\npayload:\n"
	docsArray := strings.Split(string(reqBody), "---")

	// All chunks except the first starts at line 3 for some reason... so we have to do the first one manually outside the j loop
	toLines := strings.Split(docsArray[0], "\n")
	for i = 2; i < len(toLines); i++ {
		result += toLines[i]
	}

	// Now itterate the rest of the chunks starting with index 1
	for j = 1; j < len(docsArray); j++ {
		toLines = strings.Split(docsArray[j], "\n")
		for i = 3; i < len(toLines); i++ {
			result += toLines[i]
		}
	}
	// Send result to the client
	w.Write([]byte(result))
}

func ungroup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// m[] := make(map[string]string)

	result := reqBody

	// Send result to the client
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
