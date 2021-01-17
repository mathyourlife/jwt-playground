-- As super user

DROP DATABASE IF EXISTS app_database;
DROP USER IF EXISTS app_user;
DROP ROLE IF EXISTS app_user_role;

CREATE DATABASE app_database ENCODING 'UTF8';
CREATE ROLE app_user_role LOGIN CREATEROLE;
CREATE USER app_user PASSWORD 'app_user_password';
GRANT app_user_role to app_user;
GRANT ALL PRIVILEGES ON DATABASE app_database to app_user_role;

CREATE TABLE user (
	user_id BIGSERIAL,
	username varchar(256),
	password varchar(256),
	salt varchar(256),
	CONSTRAINT user_id PRIMARY KEY(user_id)
)
