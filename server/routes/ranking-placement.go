
package routes

import (
	"net/http"
//	"database/sql"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type ranking struct {
	Id int 		`json:"id"`;
	RanCode string `json:"ran_code"`;
	Name string `json:"name"`;
}

func RankingPlacements(w http.ResponseWriter, r *http.Request) {
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
			var rankings []ranking
			for rows.Next() {
				var rank ranking
				if err := rows.Scan(
					&rank.Id,
					&rank.RanCode,
					&rank.Name,
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
