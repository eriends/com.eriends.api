package main

import (
	"fmt"
	"net/http"
)

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Eriends!")
}

