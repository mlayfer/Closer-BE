CREATE SCHEMA CloserDB;

CREATE TABLE IF NOT EXISTS CloserDB.Users (
    identifier VARCHAR(24) PRIMARY KEY,
    first_name VARCHAR(16) NOT NULL,
    last_name VARCHAR(16),
    email VARCHAR(255),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted BOOL
);

CREATE TABLE IF NOT EXISTS CloserDB.UsersPlatforms (
	platform_id  INT AUTO_INCREMENT PRIMARY KEY,
    user_identifier VARCHAR(24) NOT NULL,
    platform_type INT,
    platform_data VARCHAR(255),
    FOREIGN KEY (user_identifier)
        REFERENCES Users (identifier)
        ON DELETE RESTRICT
)