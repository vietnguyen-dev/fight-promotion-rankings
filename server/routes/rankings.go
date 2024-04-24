package routes

import (
	"net/http"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type ranking struct {
	Id int 		`json:"id"`;
	PromotionId int `json:"promotion_id"`;
	RankingPlacementId int `json:"ranking_placement_id"`;
	Ranking string `json:"ranking"`;
	WeightClassId int `json:"weight_class_id"`;
	WeightClass string `json:"weight_class"`;
	WeightFirstThreeDigits string `json:"weight_first_three_digits"`;
	AthleteId int `json:"athlete_id"`;
	AthleteName string `json:"athlete_name"`;
}

func Ranking(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			id := r.PathValue("promotions_id")
			//query the promotions view
			rows, err := db.Query("SELECT * FROM vw_rankings WHERE promotion_id = $1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var rankings []ranking
			for rows.Next() {
				var rank ranking
				if err := rows.Scan(
					&rank.Id,
					&rank.PromotionId,
					&rank.RankingPlacementId,
					&rank.Ranking,
					&rank.WeightClassId,
					&rank.WeightClass,
					&rank.WeightFirstThreeDigits,
					&rank.AthleteId,
					&rank.AthleteName,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				rankings = append(rankings, rank)
			}

			//return data
			jsonData, err := json.Marshal(rankings)
			if err != nil {
				http.Error(w, "JSON serialization failed", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}

			// Set the content type and write the JSON data
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		default:
  			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
 	}
}
