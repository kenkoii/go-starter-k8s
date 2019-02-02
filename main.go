package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := "JP Team and PH Team"
	if param := r.URL.Query().Get("name"); param != "" {
		msg = param
	}
	fmt.Fprintf(w, "Hi there, I love %s!", msg)
}

func main() {
	port := ":3000"
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
	log.Printf("Running on port %s", port)
}
