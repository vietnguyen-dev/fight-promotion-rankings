CREATE TABLE fights (
    id SERIAL PRIMARY KEY,
    athlete_id INT NOT NULL,
    athlete_two_id INT NOT NULL,
    events_id INT NOT NULL,
    weight_class_id INT NOT NULL, 
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
