package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(handleRoot))
	http.ListenAndServe(":5000", nil)
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello! (github)")
}
