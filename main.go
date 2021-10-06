package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s = %s\n", k, v))
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "ok,httpstatus:200 \n")

}

func Getenv(w http.ResponseWriter, r *http.Request) {
	for _, v := range os.Environ() {
		str := strings.Split(v, "=")
		if str[0] == "HOME" {
			io.WriteString(w, str[1])
		}

	}
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/Getenv", Getenv)
	http.ListenAndServe(":8000", nil)

}
