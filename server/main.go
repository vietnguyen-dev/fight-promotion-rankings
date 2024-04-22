package main

import (
    "fmt"
    "net/http"
)

type promotions struct {
	id int;
	pro_code string;
	name string;
	how_many_ranked int;
	website_link string;
}

// handler is a typical HTTP request-response handler in Go; details later
func handler(w http.ResponseWriter, r *http.Request) {
    // Fprintf formats the string to a writer
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/promotions", handler)
    http.ListenAndServe(":8080", nil)
}
