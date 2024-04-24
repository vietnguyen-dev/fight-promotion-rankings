package routes

import (
	"net/http"
	"encoding/json"
	"fmt"	
	"server/utils"
)

type vw_weight_classes struct {
	Id int 		`json:"id"`;
	WeiCode string `json:"wei_code"`;
	WeightRange string `json:"weight_range"`;
	WeightFirstThree string `json:"weight_first_three"`;
	Name string `json:"name"`;
	PromotionId string `json:"promotion_id"`;
}


func WeightClasses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("promotions_id")
			rows, err := db.Query("SELECT * FROM vw_weight_classes WHERE promotion_id = $1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var weight_classes []vw_weight_classes
			for rows.Next() {
				var weight_class vw_weight_classes
				if err := rows.Scan(
					&weight_class.Id,
					&weight_class.WeiCode,
					&weight_class.WeightRange,
					&weight_class.WeightFirstThree,
					&weight_class.Name,
					&weight_class.PromotionId,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				weight_classes = append(weight_classes, weight_class)
			}

			//return data
			jsonData, err := json.Marshal(weight_classes)
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
