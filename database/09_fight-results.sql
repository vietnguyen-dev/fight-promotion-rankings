CREATE TABLE fight_results (
	id SERIAL PRIMARY KEY,
	fight_id INT NOT NULL,
	winner_id INT NOT NULL,
	loser_id INT NOT NULL,
	result_id INT NOT NULL,
	event_id INT NOT NULL,
    	date_created DATE DEFAULT CURRENT_DATE,
    	date_updated DATE,
    	date_deleted DATE,
	FOREIGN KEY (fight_id) REFERENCES fights(id),
	FOREIGN KEY (winner_id) REFERENCES athletes(id),
	FOREIGN KEY (loser_id) REFERENCES athletes(id), 
	FOREIGN KEY (result_id) REFERENCES results(id),
	FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE VIEW vw_fight_results AS
SELECT
	fr.id,
	fr.fight_id,
	fr.winner_id,
	a1.last_name AS winner_last_name,
	fr.loser_id,
	a2.last_name AS loser_last_name,
	f.rounds_lasted,
	f.time_ellapsed,
	fr.result_id,
	r.name as result,
	e.event_date
FROM fight_results fr
INNER JOIN
	fights f ON f.id = fr.fight_id
INNER JOIN 
	athletes a1 ON a1.id = fr.winner_id
INNER JOIN
	athletes a2 ON a2.id = fr.loser_id
INNER JOIN
	results r ON r.id = fr.result_id
INNER JOIN
	events e ON e.id = fr.event_id
WHERE
	fr.date_deleted IS NULL
ORDER BY
	a1.promotion_id ASC;

