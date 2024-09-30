DROP DATABASE IF EXISTS records;
CREATE DATABASE records;
USE records;


CREATE TABLE Creator (
  id         INT AUTO_INCREMENT NOT NULL,
  name       VARCHAR(255) DEFAULT 'Empty',
  PRIMARY KEY (id)
);

CREATE  TABLE Sound_File (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  sound_data MEDIUMBLOB NOT NULL,
  creator_id INT NOT NULL,
  file_type VARCHAR(50) NOT NULL,
  file_size INT DEFAULT 0,
  Text_result VARCHAR(128) DEFAULT 0,
  PRIMARY KEY (id),
  FOREIGN KEY (creator_id) REFERENCES records.Creator(id)
);

INSERT INTO Creator (id, name)
    VALUES (1, 'Tester');

  CREATE USER IF NOT EXISTS 'username'@'localhost' IDENTIFIED BY 'password';
  GRANT ALL PRIVILEGES ON records.* TO 'username'@'localhost';
  FLUSH PRIVILEGES;
