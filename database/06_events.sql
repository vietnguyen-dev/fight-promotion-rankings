CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    eve_code VARCHAR(16) NOT NULL,
    event_title VARCHAR(128) NOT NULL,
    promotion_id INT NOT NULL,
    event_date DATE,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE,
    FOREIGN KEY (promotion_id) REFERENCES promotions(id)
);

CREATE VIEW vw_events AS
SELECT 
    *
FROM events
WHERE    
    date_deleted IS NULL
ORDER BY 
    promotion_id ASC, 
    COALESCE(event_date, 'infinity') DESC;
