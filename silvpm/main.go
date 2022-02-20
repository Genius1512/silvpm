package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var baseURL = "http://localhost:8000/"

func listApplications() string {
	fmt.Println("Reading application data...")
	applications, err := readJSONFormURL(baseURL + "applications/applications.json")
	if err != nil {
		fmt.Println("Error when reading valid applications; check your internet connection")
		os.Exit(1)
	}
	fmt.Println("Done.")

	for i, v := range applications {
		fmt.Printf("%v: %s\n", i+1, v)
	}
	return "list"
}

func getApplicationNameFormUser() string {
	var appName string

	// get application name
	if len(os.Args) > 1 {
		appName = os.Args[1]
	} else {
		fmt.Print("Enter application name; for a list of all applications, type 'list'> ")
		fmt.Scanln(&appName)
	}

	if appName == "list" {
		listApplications()
		return "done"
	}

	// check validity
	fmt.Println("Reading application data...")
	validAppNames, err := readJSONFormURL(baseURL + "applications/applications.json")
	if err != nil {
		fmt.Println("Error when reading valid applications; check your internet connection")
		return "error"
	}
	fmt.Println("Done.")

	if contains(validAppNames, appName) {
		return appName
	} else {
		fmt.Println("The application does not exist")
		return "error"
	}
}

func main() {
	appName := getApplicationNameFormUser()
	if appName == "error" {
		os.Exit(1)
	} else if appName == "done" {
		os.Exit(0)
	}

	fmt.Printf("Downloading %s...\n", appName)
	err := downloadApplication(appName)
	if err != nil {
		fmt.Println("Error while downloading application: ", err)
		return
	}
	fmt.Println("Done.")
}

func contains(elements []string, value string) bool {
	for _, v := range elements {
		if value == v {
			return true
		}
	}
	return false
}

func readJSONFormURL(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var jsonData []string
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}
