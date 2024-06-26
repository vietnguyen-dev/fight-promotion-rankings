CREATE TABLE rankings ( 
	id SERIAL PRIMARY KEY, 
	promotion_id INT NOT NULL,
	ranking_placement_id INT NOT NULL,
	weight_class_id INT NOT NULL,
	athlete_id INT NOT NULL, 
	FOREIGN KEY (promotion_id) REFERENCES promotions(id),
	FOREIGN KEY (ranking_placement_id) REFERENCES ranking_placement(id),
	FOREIGN KEY (weight_class_id) REFERENCES weight_classes(id), 
	FOREIGN KEY (athlete_id) REFERENCES athletes(id)

);

CREATE VIEW vw_rankings AS
SELECT
	r.id,
	r.promotion_id,
	r.ranking_placement_id,
	rp.name AS ranking,
	r.weight_class_id,
	wc.name as weight_class,
	CASE 
        	WHEN length(wc.weight_range) >= 3 THEN 
            		CAST(SUBSTRING(wc.weight_range FROM 1 FOR 3) AS INTEGER)
        	ELSE 
            		NULL
    	END AS weight_first_three_digits,
	r.athlete_id,
    	CONCAT(a.first_name, ' ', a.last_name) as athlete_name
FROM rankings r
INNER JOIN
	athletes a ON a.id = r.athlete_id
INNER JOIN
	ranking_placement rp ON rp.id = r.ranking_placement_id
INNER JOIN
	weight_classes wc on wc.id = r.weight_class_id
ORDER BY
	r.promotion_id ASC, weight_first_three_digits ASC, r.ranking_placement_id ASC;
	
	
