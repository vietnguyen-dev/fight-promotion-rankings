CREATE TABLE promotions (
    id SERIAL PRIMARY KEY,
    pro_code VARCHAR(16) NOT NULL,
    name VARCHAR(256) NOT NULL,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE
);

