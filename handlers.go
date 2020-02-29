package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func reactAppServe(w http.ResponseWriter, req *http.Request) {
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

func vegetaHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			fmt.Println(err)
			return
		}
		vo, err := mapVegetaOptions(b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			fmt.Println(err)
			return
		}
		res, ct, err := execVegetaCall(vo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			fmt.Println(err)
			return
		}
		w.Header().Add("Content-Type", ct)
		w.Write([]byte(res))
	}
}
