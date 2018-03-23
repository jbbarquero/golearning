package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plai")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	fmt.Println("HTTPS...")

	http.HandleFunc("/", handler)

	go http.ListenAndServe(":8080", http.RedirectHandler("https://localhost:10443", 301))

	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)

	fmt.Println("HTTPS.END .")
}
