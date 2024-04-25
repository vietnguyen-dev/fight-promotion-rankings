package routes

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type vw_promotions struct {
	Id int 		`json:"id"`;
	ProCode string `json:"pro_code"`;
	Name string `json:"name"`;
	Email string `json:"email"`;
	HowManyRanked int `json:"how_many_ranked"`;
	WebsiteLink string `json:"website_link"`;
	S3Url string `json:"s3_url"`;
}


func Promotions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("id")
			rows, err := db.Query("SELECT * FROM vw_promotions WHERE id = $1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var promotions []vw_promotions
			for rows.Next() {
				var promo vw_promotions
				var websiteLink sql.NullString
				var s3url sql.NullString 
				if err := rows.Scan(
					&promo.Id,
					&promo.ProCode,
					&promo.Name,
					&promo.Email,
					&promo.HowManyRanked,
					&websiteLink,
					&s3url,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				promo.WebsiteLink = utils.NullStringToString(websiteLink)
				promo.S3Url = utils.NullStringToString(s3url)
				promotions = append(promotions, promo)
			}

			//return data
			jsonData, err := json.Marshal(promotions)
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
