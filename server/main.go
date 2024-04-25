package main

import ( 
	"net/http"
	"server/routes"
)

func main() {
	http.HandleFunc("GET /promotions/{id}", routes.Promotions)

	//athletes routes
	http.HandleFunc("GET /athletes/", routes.GetAthletes)

	//results route
	http.HandleFunc("/results", routes.Results)
	http.HandleFunc("/events/{promotions_id}", routes.Events)
	http.HandleFunc("/weight-classes/{promotions_id}", routes.WeightClasses)
	http.HandleFunc("/fights/{events_id}", routes.Fights)
	http.HandleFunc("/fight-results/{fight_id}", routes.FightResults)
	http.HandleFunc("/ranking-placements", routes.RankingPlacements)
	http.HandleFunc("/ranking/{promotions_id}", routes.Ranking)
	http.ListenAndServe(":8080", nil)
}
