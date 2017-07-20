package mlclassifier

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXIST mlclassifier DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE mlclassifier;`
	`CREATE TABLE users (
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    email TEXT NOT NULL,
	    password TEXT NOT NULL,
	    token TEXT NOT NULL,
	    modelLimit INT UNSIGNED NOT NULL,
	    createdAt DATETIME NOT NULL,
	    lastLoginAt DATETIME NOT NULL,
	    PRIMARY KEY (id)
	);`,
	`CREATE TABLE models (
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    idUser INT UNSIGNED NOT NULL,
	    url TEXT NOT NULL,
	    trainingLimit INT UNSIGNED NOT NULL,
	    requestLimit INT UNSIGNED NOT NULL,
	    createdAt DATETIME NOT NULL,
	    PRIMARY KEY (id),
	    FOREIGN KEY (idUser) REFERENCES users(id)
	);`,
	`CREATE TABLE trainings (
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    idModel INT NOT NULL,
	    time INT UNSIGNED NOT NULL,
	    createdAt DATETIME NOT NULL,
	    PRIMARY KEY (id),
	    FOREIGN KEY (idModel) REFERENCES models(id)
	);`,
	`CREATE TABLE request (
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    idModel INT UNSIGNED NOT NULL,
	    frequency TEXT NOT NULL,
	    hits INT UNSIGNED NOT NULL,
	    avgProcessingTime INT UNSIGNED NOT NULL,
	    createdAt DATETIME NOT NULL,
	    PRIMARY KEY (id),
	    FOREIGN KEY (idModel) REFERENCES models(id)
	);`,
}