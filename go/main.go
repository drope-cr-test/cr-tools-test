package main

import (
	"fmt"
	"net/http"
)

func main() {
	// govet printf canary: %d with a string argument.
	fmt.Printf("CR-MARKER-T6-VET: %d\n", "not-a-number")
	// staticcheck SA4006 canary: first value of x is never used.
	x := 1
	x = 2
	fmt.Println("hello")
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
