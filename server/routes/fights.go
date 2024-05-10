package routes

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type vw_fights struct {
	Id int 		`json:"id"`;
	AthleteId int `json:"athlete_id"`;
	AthleteName string `json:"athlete_name"`;
	AthleteTwoId int `json:"athlete_two_id"`;
	AthleteTwoName string `json:"athlete_two_name"`;
	EventsId int `json:"events_id"`;
	PromotionId int `json:"promotion_id"`;
	WeightClassId int `json:"weight_class_id"`;
	WeightClassName string `json:"weight_class_name"`;
	CardPlace int `json:"card_place"`;
	TimeEllapsed string `json:"time_ellapsed"`;
	RoundsLasted int `json:"rounds_lasted"`;
	NumberRounds int `json:"number_rounds"`;
	MinutesPerRound int `json:"min_per_round"`;
}


func Fights(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("events_id")
			rows, err := db.Query("SELECT * FROM vw_fights WHERE events_id = $1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var fights []vw_fights
			for rows.Next() {
				var fight vw_fights
				var athleteId sql.NullInt64
				var athleteName sql.NullString
				var athleteTwoId sql.NullInt64
				var athleteTwoName sql.NullString
				var weightClassId sql.NullInt64
				var weightClassName sql.NullString
				var cardPlace sql.NullInt64
				var timeEllapsed sql.NullString
				var roundsLasted sql.NullInt64
				var numberRounds sql.NullInt64
				var minutesPerRound sql.NullInt64
				
				if err := rows.Scan(
					&fight.Id,
					&athleteId,
					&athleteName,
					&athleteTwoId,
					&athleteTwoName,
					&fight.PromotionId,
					&fight.EventsId,
					&weightClassId,
					&weightClassName,
					&cardPlace,
					&timeEllapsed,
					&roundsLasted,
					&numberRounds,
					&minutesPerRound,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				fight.AthleteId = utils.NullInt64ToInt(athleteId)
				fight.AthleteName = utils.NullStringToString(athleteName)
				fight.AthleteTwoId = utils.NullInt64ToInt(athleteTwoId)
				fight.AthleteTwoName = utils.NullStringToString(athleteTwoName)
				fight.WeightClassId = utils.NullInt64ToInt(weightClassId)
				fight.WeightClassName = utils.NullStringToString(weightClassName)
				fight.CardPlace = utils.NullInt64ToInt(cardPlace)
				fight.TimeEllapsed = utils.NullStringToString(timeEllapsed)
				fight.RoundsLasted = utils.NullInt64ToInt(roundsLasted)
				fight.NumberRounds = utils.NullInt64ToInt(numberRounds)
				fight.MinutesPerRound = utils.NullInt64ToInt(minutesPerRound)

				fights = append(fights, fight)
			}

			//return data
			jsonData, err := json.Marshal(fights)
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
