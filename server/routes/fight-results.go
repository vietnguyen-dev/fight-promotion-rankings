package routes

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"	
	"time"
	"server/utils"
)

type vw_fight_results struct {
	Id int 		`json:"id"`;
	FightId int `json:"fight_id"`;
	WinnerId int `json:"winner_id"`;
	WinnerLastName string `json:"winner_last_name"`;
	LoserId int `json:"loser_id"`;
	LoserLastName string `json:"loser_last_name"`;
	RoundsLasted int `json:"rounds_lasted"`;
	TimeEllapsed string `json:"time_ellapsed"`;
	ResultId int `json:"result_id"`;
	Result string `json:"result"`;
	EventDate time.Time `json:"event_date"`;
}


func FightResults(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("fight_id")
			rows, err := db.Query("SELECT * FROM vw_fight_results WHERE fight_id = $1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var fight_results []vw_fight_results
			for rows.Next() {
				var fight_result vw_fight_results
				var timeEllapsed sql.NullString
				var roundsLasted sql.NullInt64
				var eventDate sql.NullTime
				
				if err := rows.Scan(
					&fight_result.Id,
					&fight_result.FightId,
					&fight_result.WinnerId,
					&fight_result.WinnerLastName,
					&fight_result.LoserId,
					&fight_result.LoserLastName,
					&roundsLasted,
					&timeEllapsed,
					&fight_result.ResultId,
					&fight_result.Result,
					&eventDate,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				fight_result.RoundsLasted = utils.NullInt64ToInt(roundsLasted)
				fight_result.TimeEllapsed = utils.NullStringToString(timeEllapsed)
				fight_result.EventDate = utils.NullTimeToTime(eventDate)

				fight_results = append(fight_results, fight_result)
			}

			//return data
			jsonData, err := json.Marshal(fight_results)
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
