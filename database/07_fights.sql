CREATE TABLE fights (
    id SERIAL PRIMARY KEY,
    athlete_id INT,
    athlete_two_id INT,
    events_id INT NOT NULL,
    weight_class_id INT,  
    card_place INT,
    time_ellapsed VARCHAR(8),
    rounds_lasted INT,
    number_rounds INT,
    min_per_round INT,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE,
    FOREIGN KEY (events_id) REFERENCES events(id),
    FOREIGN KEY (weight_class_id) REFERENCES weight_classes(id),
    FOREIGN KEY (athlete_id) REFERENCES athletes(id),
    FOREIGN KEY (athlete_two_id) REFERENCES athletes(id)
);

CREATE VIEW vw_fights AS
SELECT
    f.id,
    f.athlete_id,
    CONCAT(a1.first_name, ' ', a1.last_name) as athlete_name,
    f.athlete_two_id,
    CONCAT(a2.first_name, ' ', a2.last_name) as athlete_two_name,
    f.events_id,
    e.promotion_id AS promotion_id,
    f.weight_class_id,
    wc.name as weight_class_name,
    f.card_place,
    f.time_ellapsed,
    f.rounds_lasted,
    f.number_rounds,
    f.min_per_round
FROM fights f
INNER JOIN 
	events e ON e.id = f.events_id
INNER JOIN
	weight_classes wc ON wc.id = f.weight_class_id
LEFT JOIN
	athletes a1 ON a1.id = f.athlete_id
LEFT JOIN
	athletes a2 on a2.id = f.athlete_two_id
WHERE 
	f.date_deleted IS NULL
ORDER BY 
	e.promotion_id ASC, events_id ASC, f.card_place NULLS LAST;
