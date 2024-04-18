CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    eve_code VARCHAR(16) NOT NULL,
    event_title VARCHAR(128) NOT NULL,
    pid INT NOT NULL,
    event_date DATE,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE,
    FOREIGN KEY (pid) REFERENCES promotions(id)
);
