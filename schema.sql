DROP DATABASE IF EXISTS bank_app;
CREATE DATABASE bank_app;
USE bank_app;
CREATE TABLE users (
    id VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (email)
);