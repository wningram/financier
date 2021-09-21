CREATE DATABASE financier;

\c financier;

CREATE TABLE Users (
    Username varchar(30),
    FirstName text,
    LastName text
);

INSERT INTO Users VALUES ('admin', 'Admin', 'User');