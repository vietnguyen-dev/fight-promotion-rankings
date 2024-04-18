CREATE TABLE rankings ( 
	id SERIAL PRIMARY KEY, 
	promotion_id INT NOT NULL,
	rank_placement_id INT NOT NULL,
	weight_class_id INT NOT NULL,
	athlete_id INT NOT NULL, 
	FOREIGN KEY (promotion_id) REFERENCES promotions(id),
	FOREIGN KEY (rank_placement_id) REFERENCES ranking_placement(id),
	FOREIGN KEY (weight_class_id) REFERENCES weight_classes(id), 
	FOREIGN KEY (athlete_id) REFERENCES athletes(id)

);
