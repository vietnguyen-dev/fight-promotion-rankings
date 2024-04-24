package routes

import (
	"net/http"
//	"database/sql"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type tresults struct {
	Id int 		`json:"id"`;
	ResCode string `json:"res_code"`;
	Name string `json:"name"`;
}

func Results(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			rows, err := db.Query("SELECT * FROM results;")
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var results []tresults
			for rows.Next() {
				var res tresults
				if err := rows.Scan(
					&res.Id,
					&res.ResCode,
					&res.Name,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				results = append(results, res)
			}

			//return data
			jsonData, err := json.Marshal(results)
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
