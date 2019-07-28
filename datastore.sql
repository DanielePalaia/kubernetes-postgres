DROP DATABASE tododatastore;

CREATE DATABASE tododatastore;

 \c tododatastore

CREATE TABLE ToDo (
	    ID serial,
	    Topic varchar(255),
	    Completed int,
	    Due varchar(255) DEFAULT '',
	    PRIMARY KEY (ID)
);
