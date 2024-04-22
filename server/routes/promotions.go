package promotionsHandler

import (
    "fmt"
    "net/http"
)

func promotionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Contact us via GET request.")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "Contact us via POST request.")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

