DROP TABLE IF EXISTS album;
DROP TABLE IF EXISTS album;
DROP TABLE IF EXISTS album;

CREATE DATABASE records
CHARACTER SET = utf8mb4
COLLATE = utf8mb4_general_ci;

USE records;


CREATE TABLE Creator (
  id         INT AUTO_INCREMENT NOT NULL,
  name       VARCHAR(255) DEFAULT 'Empty',
  PRIMARY KEY (`id`)
);

CREATE  TABLE Sound_File (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  sound_data MEDIUMBLOB NOT NULL,
  creator_id INT NOT NULL,
  file_path VARCHAR(255),
  file_size INT,
  Text_result VAR(128),
  FOREIGN KEY (creator_id) REFERENCES records.Creator(id)
)



  CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
  GRANT ALL PRIVILEGES ON records.* TO 'username'@'localhost';
  FLUSH PRIVILEGES;
