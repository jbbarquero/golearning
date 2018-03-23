package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("Webserver...")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Webserver. End.")
}
