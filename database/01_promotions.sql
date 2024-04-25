CREATE TABLE promotions (
    id SERIAL PRIMARY KEY,
    pro_code VARCHAR(16) NOT NULL,
    name VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL,
    how_many_ranked INT NOT NULL,
    website_link VARCHAR,
    s3_url VARCHAR,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE
);

CREATE VIEW vw_promotions AS
SELECT 
	id,
	pro_code,
	name,
	email,
	how_many_ranked,
	website_link,
	s3_url
FROM promotions
WHERE
	date_deleted IS NULL;
