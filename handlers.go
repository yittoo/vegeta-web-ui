package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func reactAppProxy(w http.ResponseWriter, req *http.Request) {
	if !isDevelopment {
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
	} else {
		fmt.Fprintf(w, "This is development env, please view React app itself directly for debugging")
	}
}

func vegeta(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
