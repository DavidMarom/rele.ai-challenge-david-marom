package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func ungroup(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hey")
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

func group(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	i := 0

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Break to lines and remove first 2
	linesArray := strings.Split(string(reqBody), "\n")
	// remove first 2 indexes - deep copy
	arrayMinusTwoFirstLines := make([]string, len(linesArray))
	for i = 2; i < len(linesArray); i++ {
		arrayMinusTwoFirstLines[i-2] = linesArray[i]
	}

	// seperate by ":" and insert to map

	myMap := make(map[string]string)
	tmpString := ""

	for i = 0; i < len(arrayMinusTwoFirstLines)-2; i++ {
		tmpString = arrayMinusTwoFirstLines[i]

		// insert a[0] and a[1] to map in a form of key:value pairs
		myMap[strings.Split(tmpString, ":")[0]] = strings.Split(tmpString, ":")[1]
	}

	// Prepare the final string with header
	result := "Line 01\nLine 02\n"

	// itterate over config file and add to result

	if myMap["aaa"] != "" {
		result += "aaa:" + myMap["aaa"]
	}

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
}
