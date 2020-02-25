package main

import (
	"net/http"
	"os"
)

var isDevelopment = false

func init() {
	args := os.Args
	// add all possible runtime arguments here
	for _, v := range args {
		switch v {
		case "dev":
			isDevelopment = true
			break
		}
	}
}

func main() {

	if !isDevelopment {
		startReactApp()
		http.HandleFunc("/", reactAppProxy)
	}
	http.HandleFunc("/vegeta", vegeta)
	http.ListenAndServe(":8000", nil)
}
