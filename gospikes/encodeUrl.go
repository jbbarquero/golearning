package main

import (
	"fmt"
	"net/url"
)

func main() {

	path := "/path"

	encodedPath := url.PathEscape(path)

	fmt.Printf("Encode %s to %s. Done.\n", path, encodedPath)

	fmt.Printf("Again, encode %s to %s. Done.\n", path, url.QueryEscape(path))
}
