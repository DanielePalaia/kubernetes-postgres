DROP DATABASE tododatastore;

CREATE DATABASE tododatastore;

USE tododatastore

CREATE TABLE ToDo (
	    ID int NOT NULL AUTO_INCREMENT,
	    Topic varchar(255),
	    Completed int,
	    Due varchar(255) DEFAULT '',
	    PRIMARY KEY (ID)
);
