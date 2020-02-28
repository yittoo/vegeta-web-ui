package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// TODO serve static files properly
func reactAppProxy(w http.ResponseWriter, req *http.Request) {
	if !isDevelopment {
		f, err := os.Open("./client/build/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened"))
			fmt.Println(err)
			return
		}
		defer f.Close()

		s := bufio.NewScanner(f)
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
