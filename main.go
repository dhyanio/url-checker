package main

// This simple URL Checking Golang Program.

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestStatus = errors.New("this is a status code is failed ")

func main() {
	fmt.Printf("this is a url checker golag program\n\n")

	urls := []string{
		"https://www.google.com",
		"https://www.amazon.in",
		"https://www.flipkart.com",
	}

	for _, url := range urls {
		hitURL(url)

	}
}

// hitUrl is a url Hit function
func hitURL(url string) error {
	fmt.Println("checking this : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestStatus
	}
	return nil

}
