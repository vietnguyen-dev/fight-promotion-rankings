package routes

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"	
	"time"
	"server/utils"
)

type vw_athletes struct {
	Id		 int 		`json:"id"`;
	FirstName 	 string `json:"first_name"`;
	LastName 	 string `json:"last_name"`;
	NickName 	 string `json:"nickname"`;
	GymName 	 string `json:"gym_name"`;
	Location 	 string `json:"location"`;
	Height 		 string `json:"height"`;	
	Reach 		 int `json:"reach"`;
	Dob 		 time.Time `json:"dob"`;
	ActiveStatus 	 bool `json:"active_status"`;
	PromotionId 	 int `json:"promotion_id"`;
	InstagramLink 	 string `json:"instagram_link"`;
	TwitterLink 	 string `json:"twitter_link"`;
	FacebookLink 	 string `json:"facebook_link"`;
	S3Url 		 string `json:"s3_url"`;
}

func Athletes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("promotions_id")
			rows, err := db.Query("SELECT * FROM vw_athletes WHERE promotion_id = $1 LIMIT 1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var athletes []vw_athletes
			for rows.Next() {
				var athlete vw_athletes
				var NickName sql.NullString
				var GymName sql.NullString
				var Height sql.NullString
				var Location sql.NullString
				var Reach sql.NullInt64
				var Dob sql.NullTime
				var ActiveStatus sql.NullBool
				var InstagramLink sql.NullString
				var TwitterLink sql.NullString
				var FacebookLink sql.NullString
				var S3Url sql.NullString
				if err := rows.Scan(
					&athlete.Id,
					&athlete.FirstName,
					&athlete.LastName,
					&NickName,
					&GymName,
					&Location,
					&Height,
					&Reach,
					&Dob,
					&ActiveStatus,
					&athlete.PromotionId,
					&InstagramLink,
					&TwitterLink,
					&FacebookLink,
					&S3Url,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				athlete.NickName = utils.NullStringToString(NickName) 
				athlete.GymName = utils.NullStringToString(GymName)
				athlete.Height = utils.NullStringToString(Height)
				athlete.Location = utils.NullStringToString(Location)
				athlete.Reach = utils.NullInt64ToInt(Reach)
				athlete.Dob = utils.NullTimeToTime(Dob)
				athlete.ActiveStatus= utils.NullBoolToBool(ActiveStatus) 
				athlete.InstagramLink = utils.NullStringToString(InstagramLink)
				athlete.TwitterLink = utils.NullStringToString(TwitterLink)
				athlete.FacebookLink = utils.NullStringToString(FacebookLink)
				athlete.S3Url = utils.NullStringToString(S3Url)
				athletes = append(athletes, athlete)
			}

			//return data
			jsonData, err := json.Marshal(athletes)
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
