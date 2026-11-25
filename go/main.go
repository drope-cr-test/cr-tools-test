package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// govet printf canary: %d with a string argument.
	fmt.Printf("CR-MARKER-T6-VET: %d\n", "not-a-number")
	// staticcheck SA4006 canary: first value of x is overwritten before any read.
	// Nonconstant initializer (len(os.Args)) so SA4006 does not skip it.
	x := len(os.Args)
	x = 2
	fmt.Println(x)
	// bodyclose canary: response body never closed.
	resp, _ := http.Get("https://example.com")
	fmt.Println(resp.Status)
}

func getThing() (*http.Response, error) {
	resp, err := http.Get("https://example.com")
	if err != nil {
		// nilerr canary: returning nil error alongside a nil value.
		return nil, nil
	}
	return resp, nil
}
