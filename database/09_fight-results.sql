CREATE TABLE fight_results (
	id SERIAL PRIMARY KEY,
	fight_id INT NOT NULL,
	winner_id INT NOT NULL,
	loser_id INT NOT NULL,
	result_id INT NOT NULL,
    	date_created DATE DEFAULT CURRENT_DATE,
    	date_updated DATE,
    	date_deleted DATE,
	FOREIGN KEY (fight_id) REFERENCES fights(id),
	FOREIGN KEY (winner_id) REFERENCES athletes(id),
	FOREIGN KEY (loser_id) REFERENCES athletes(id), 
	FOREIGN KEY (result_id) REFERENCES results(id)
);
