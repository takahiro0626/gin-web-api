CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    mail VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);