CREATE TABLE category (
	id serial PRIMARY KEY,
	title VARCHAR ( 50 ) UNIQUE NOT NULL,
	description VARCHAR ( 50 ),
	createddate TIMESTAMP,
    updatedate TIMESTAMP
);