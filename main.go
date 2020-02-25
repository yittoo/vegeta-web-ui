package main

import (
	"fmt"
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
	}
	fmt.Println("Hellooo World!")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8000", nil)
}
