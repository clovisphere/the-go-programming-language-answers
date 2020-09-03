package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// add http:// prefix if it's not present
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// NOTE: only interested in the HTTP status code returned,
		// disregarding resp.Body content
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		fmt.Printf("Total number of bytes: %v\nWhat was the HTTP status code returned? %v\n", nbytes, resp.Status)
	}
}
