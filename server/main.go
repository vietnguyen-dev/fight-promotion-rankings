package main

import ( 
	"net/http"
	"server/routes"
)

func main() {
	http.HandleFunc("/promotions/{id}", routes.Promotions)
	http.HandleFunc("/athletes/{promotions_id}", routes.Athletes)
	http.HandleFunc("/results", routes.Results)
	http.ListenAndServe(":8080", nil)
}
