package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func gorilla_handler(rw http.ResponseWriter, req *http.Request) {
	val := req.URL.Path[2:]
	if val == "yes" {
		fmt.Fprintf(rw, "So you like gorillas, huh?")
	} else if val == "no" {
		fmt.Fprintf(rw, "Why don't you like gorillas, huh?")
	}
}

func main() {
	http.HandleFunc("/gorillas", gorilla_handler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
