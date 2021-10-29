CREATE TABLE post (
	id serial4 NOT NULL,
	title varchar(50) NOT NULL,
	body text NULL,
	createdate timestamp NULL,
	updatedate timestamp NULL,
	isshare boolean NULL,
	isactive boolean NULL,
	categoryid numeric NULL
);