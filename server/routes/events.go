package routes

import (
	"net/http"
	"encoding/json"
	"fmt"	
	"time"
	"server/utils"
)

type vw_events struct {
	Id int 		`json:"id"`;
	EveCode string `json:"eve_code"`;
	EventTitle string `json:"event_title"`;
	PromotionId int `json:"promotion_id"`;
	EventDate time.Time `json:"website_link"`;
}


func Events(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("promotions_id")
			rows, err := db.Query("SELECT * FROM vw_events WHERE id = $1 LIMIT 1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var events []vw_events
			for rows.Next() {
				var event vw_events
				if err := rows.Scan(
					&event.Id,
					&event.EveCode,
					&event.EventTitle,
					&event.PromotionId,
					&event.EventDate,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				events = append(events, event)
			}

			//return data
			jsonData, err := json.Marshal(events)
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
