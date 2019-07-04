package main

import (
	"fmt"
	"net/http"
	"time"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 20)
	fmt.Fprintf(w, "Hello from the very best server!")
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Running")
}
