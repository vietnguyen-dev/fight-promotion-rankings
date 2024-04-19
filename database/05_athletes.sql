CREATE TABLE athletes (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256) NOT NULL,
    nickname VARCHAR(128),
    gym_name VARCHAR(128),
    location VARCHAR(256),
    height VARCHAR(8),
    reach INT,
    dob DATE,
    active_status BOOL,
    promotion_id INT NOT NULL,
    instagram_link VARCHAR(255),
    twitter_link VARCHAR(255),
    facebook_link VARCHAR(255),
    s3_url VARCHAR,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE,
    FOREIGN KEY (promotion_id) REFERENCES promotions(id)
);

CREATE VIEW vw_athletes AS
SELECT
	*
FROM athletes
WHERE 
	date_deleted IS NULL
ORDER BY 
	promotion_id ASC;
	

