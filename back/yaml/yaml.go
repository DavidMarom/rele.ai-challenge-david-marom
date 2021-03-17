package yaml

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

// ******  UNGROUP ROUTE  ******
func Ungroup(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	j := 0
	result := "type: DataFile\npayload:\n"
	docsArray := strings.Split(string(reqBody), "---")

	// All chunks except the first starts at line 3 so we have to do the first one manually outside the j loop below
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

// ******  GROUP ROUTE ******
func Group(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	linesArray := strings.Split(string(reqBody), "\n")         // Break to lines and remove first 2
	arrayMinusTwoFirstLines := make([]string, len(linesArray)) // remove first 2 indexes - deep copy

	i := 0
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

	result := "" // Prepare the final string

	// *** itterate over config file and add to result ***

	confFile, err := ioutil.ReadFile("group-config.yaml") // read config.YAML file...
	if err != nil {
		fmt.Print(err)
	}
	confFileArr := strings.Split(string(confFile), "\n") // ... and put it in an array

	// gush = key : title
	lastItem := ""
	isFirstTime := true // the --- shouldn't appear before the first group so we'll use this var as a flag

	for i, gush := range confFileArr { // itterate over the config array

		// if myMap contains a key like gush[0] - add the title and the map key:value
		if myMap[strings.Split(gush, ":")[0]] != "" {
			if lastItem != strings.Split(gush, ":")[1] {
				if !isFirstTime {
					result += "---\n" // we add this only if it's not the first time
				} else {
					isFirstTime = false
				}
				result += strings.Split(gush, ":")[1] + "\n" + "payload:\n"
			}
			result += "\t" + strings.Split(gush, ":")[0] + ":" + myMap[strings.Split(gush, ":")[0]]
			lastItem = strings.Split(gush, ":")[1]
		}
		fmt.Println(i)
	}

	// Send result to the client
	w.Write([]byte(result))
}
