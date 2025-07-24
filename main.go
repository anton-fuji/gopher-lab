package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Hi, %q!\n", name)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", hiHandler)
	http.ListenAndServe(":8000", nil)
}
