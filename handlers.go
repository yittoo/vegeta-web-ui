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
		http.Redirect(w, req, "http://localhost:3000", http.StatusSeeOther)
		fmt.Printf("\n[+] This is development env, please view React app itself directly for debugging\n")
		fmt.Printf("[+] Redirecting to http://localhost:3000\n")
	}
}

func vegetaHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
