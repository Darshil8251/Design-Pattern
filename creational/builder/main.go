package main

import (
	"design-pattern/creational/builder/http"
	"fmt"
)

func main() {

	fmt.Println("Builder Pattern")
	// Create a new HTTP client
	client := http.CreateHTTPClient(http.WithAuth("Bearer token"), http.WithRateLimit(100))
	fmt.Printf("HTTP Client: %+v\n", client)

}
