package main

import ( 
	"net/http"
	"server/routes"
)

func main() {
	http.HandleFunc("/promotions/{id}", routes.Promotions)
	http.HandleFunc("/athletes/{promotions_id}", routes.Athletes)
	http.HandleFunc("/results", routes.Results)
	http.HandleFunc("/events/{promotions_id}", routes.Events)
	http.HandleFunc("/weight-classes/{promotions_id}", routes.WeightClasses)
	http.HandleFunc("/fights/{events_id}", routes.Fights)
	http.HandleFunc("/fight-results/{fight_id}", routes.FightResults)
	http.ListenAndServe(":8080", nil)
}
