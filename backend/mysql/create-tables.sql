DROP DATABASE IF EXISTS records;
CREATE DATABASE records;
USE records;


CREATE TABLE Creators (
  id         INT AUTO_INCREMENT NOT NULL,
  name       VARCHAR(255) DEFAULT 'Empty',
  PRIMARY KEY (id)
);

CREATE  TABLE Sound_Files (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  file_path VARCHAR(50) NOT NULL,
  creator_id INT NOT NULL ,
  file_type VARCHAR(50) NOT NULL,
  file_size INT DEFAULT 0,
  text_result VARCHAR(128) DEFAULT 0,
  PRIMARY KEY (id),
  FOREIGN KEY (creator_id) REFERENCES records.Creators(id)
);

INSERT INTO Creators (id, name)
    VALUES (1, 'Tester');

  CREATE USER IF NOT EXISTS 'username'@'localhost' IDENTIFIED BY 'password';
  GRANT ALL PRIVILEGES ON records.* TO 'username'@'localhost';
  FLUSH PRIVILEGES;
