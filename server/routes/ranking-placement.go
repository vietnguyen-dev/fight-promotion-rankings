package routes

import (
	"net/http"
//	"database/sql"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type tranking_placement struct {
	Id int 		`json:"id"`;
	RanCode string `json:"ran_code"`;
	Name string `json:"name"`;
}

func RankingPlacements(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			rows, err := db.Query("SELECT * FROM ranking_placement;")
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var ranking_placements []tranking_placement
			for rows.Next() {
				var ranking_placement tranking_placement
				if err := rows.Scan(
					&ranking_placement.Id,
					&ranking_placement.RanCode,
					&ranking_placement.Name,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				ranking_placements = append(ranking_placements, ranking_placement)
			}

			//return data
			jsonData, err := json.Marshal(ranking_placements)
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
