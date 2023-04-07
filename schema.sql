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
CREATE TABLE user_roles (
    id VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE TABLE bank_accounts (
    id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    balance DECIMAL(10,2) DEFAULT 0,
    currency VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);