package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/icza/gox/osx"
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
	port := 8000

	if !isDevelopment {
		buildReactApp()
		r := fmt.Sprintf("http://localhost:%v", port)
		osx.OpenDefault(r)
	}
	http.HandleFunc("/", reactAppServe)
	http.HandleFunc("/vegeta", vegetaHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static"))))
	fmt.Printf("\n[+] Serving API at port %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
