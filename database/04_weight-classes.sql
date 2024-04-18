CREATE TABLE weight_classes (
    id SERIAL PRIMARY KEY,
    wei_code VARCHAR(16) NOT NULL,
    weight_range VARCHAR(16) NOT NULL,
    name VARCHAR(256) NOT NULL,
    pid INT NOT NULL,
    date_created DATE DEFAULT CURRENT_DATE,
    date_updated DATE,
    date_deleted DATE,
    FOREIGN KEY (pid) REFERENCES promotions(id)
);

