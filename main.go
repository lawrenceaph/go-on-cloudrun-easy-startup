package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Dockerized Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
}
