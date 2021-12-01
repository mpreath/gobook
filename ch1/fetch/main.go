package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		// Exercise 1.8 - add http:// prefix if the url doesn't have one
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Exercise 1.7 - use io.Copy instead of ioutil.ReadAll
		b, err := io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// Exercise 1.9 - print the HTTP status code
		fmt.Printf("[Status code: %d]\t[Bytes read: %d]\n", resp.StatusCode, b)
	}
}
