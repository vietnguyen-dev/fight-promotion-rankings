CREATE TABLE promotions (
    id SERIAL PRIMARY KEY,
    pro_code VARCHAR(16) NOT NULL,
    name VARCHAR(256) NOT NULL,
    how_many_ranked INT NOT NULL,
    website_link VARCHAR,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE
);

CREATE VIEW vw_promotions AS
SELECT 
	*
FROM promotions
WHERE
	date_deleted IS NULL;
