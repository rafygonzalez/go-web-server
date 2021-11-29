package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Handler!")
}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Auth!")
}