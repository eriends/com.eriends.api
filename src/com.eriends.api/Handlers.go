package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Eriends!")
}

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, logo)
	fmt.Fprintf(w, "Eriends Api")
}

func ApiId(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Eriends Api Id")
}

