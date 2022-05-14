CREATE DATABASE transactional;
USE transactional;

DROP TABLE IF EXISTS user;
CREATE TABLE user
(
    id   varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    PRIMARY KEY (id)
);
SELECT * FROM user;
# TRUNCATE TABLE user;

DROP TABLE IF EXISTS user_longest_name;
CREATE TABLE user_longest_name
(
    id   varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    PRIMARY KEY (id)
);
SELECT * FROM user_longest_name;
# TRUNCATE TABLE user_longest_name;