package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, desenvol-v")
}