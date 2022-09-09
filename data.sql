CREATE DATABASE Staycation;
USE Staycation;

DROP TABLE IF EXISTS StaycationUsers;
CREATE TABLE StaycationUsers (
                                 id int NOT NULL AUTO_INCREMENT,
                                 name varchar(100) NOT NULL,
                                 dateOfBirth date NOT NULL,
                                 city varchar(100) NOT NULL,
                                 zipcode varchar(10) NOT NULL,
                                 status tinyint(1) NOT NULL DEFAULT '1',
                                 PRIMARY KEY (id)
);

INSERT INTO StaycationUsers (name, dateOfBirth, city, zipcode, status) VALUES
                                                                           ("John", "1978-12-15", "Zagreb", "10000", 1),
                                                                           ("Doe", "2000-11-25", "Berlin", "10001", 1),
                                                                           ("Johhan", "1978-10-17", "Split", "22000", 0),
                                                                           ("Sebastian", "1978-11-12", "Barcelona", "45678", 0),
                                                                           ("Bach", "1978-12-30", "Rome", "12345", 1);

CREATE TABLE Accounts (
                                 id int NOT NULL AUTO_INCREMENT,
                                 userId int NOT NULL,
                                 openingDate date NOT NULL,
                                 accountType varchar(20) NOT NULL,
                                 amount decimal NOT NULL,
                                 status tinyint(1) NOT NULL DEFAULT '1',
                                 PRIMARY KEY (id),
                                 FOREIGN KEY (userId) REFERENCES StaycationUsers(id)
);

CREATE TABLE Users (
                          id int NOT NULL AUTO_INCREMENT,
                          userId int,
                          username VARCHAR(100) NOT NULL,
                          password VARCHAR(100) NOT NULL,
                          role VARCHAR(10),
                          PRIMARY KEY (id),
                          FOREIGN KEY (userId) REFERENCES StaycationUsers(id)
);