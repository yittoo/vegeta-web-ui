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
	http.HandleFunc("/", reactAppProxy)
	http.HandleFunc("/vegeta", vegeta)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static"))))
	port := 8000
	fmt.Printf("\n[+] Serving API at port %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
