CREATE TABLE weight_classes ( 
	id SERIAL PRIMARY KEY, 
	wei_code VARCHAR(16) NOT NULL, 
	weight_range VARCHAR(16) NOT NULL,
	name VARCHAR(256) NOT NULL,
	promotion_id INT NOT NULL, 
	date_created DATE DEFAULT CURRENT_DATE,
    	date_updated DATE,
    	date_deleted DATE,
    	FOREIGN KEY (promotion_id) REFERENCES promotions(id)
);

CREATE VIEW vw_weight_classes AS
SELECT
	*,
	CASE 
        	WHEN length(weight_range) >= 3 THEN 
            		CAST(SUBSTRING(weight_range FROM 1 FOR 3) AS INTEGER)
        	ELSE 
            		NULL
    	END AS weight_first_three_digits
FROM weight_classes
WHERE 
	date_deleted IS NULL
ORDER BY 
	promotion_id ASC, weight_first_three_digits ASC;

