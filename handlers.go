package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func reactAppProxy(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Fprintf(w, "Client server is not running at the moment")
		return
	}
	defer resp.Body.Close()

	s := bufio.NewScanner(resp.Body)
	var b bytes.Buffer
	for s.Scan() {
		b.WriteString(s.Text())
	}
	fmt.Fprintf(w, b.String())
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
