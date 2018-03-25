package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTTPS Client...")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			// Get https://localhost:10443/: x509: certificate signed by unknown authority if false
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	url := "https://localhost:10443/"
	res, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error accesing %s, error: %s", url, err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body, error: %s", err)
	}
	fmt.Printf("%v\n", res.Status)
	fmt.Printf(string(body))

	fmt.Println("HTTPS Client. END.")
}
