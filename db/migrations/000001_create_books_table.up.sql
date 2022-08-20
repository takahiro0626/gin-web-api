CREATE TABLE IF NOT EXISTS books(
    id serial PRIMARY KEY,
    title VARCHAR (256) NOT NULL,
    price INT NOT NULL
);