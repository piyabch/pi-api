CREATE DATABASE IF NOT EXISTS pi_data;
USE pi_data;

CREATE TABLE IF NOT EXISTS user (
  id         INT AUTO_INCREMENT NOT NULL,
  firstname  VARCHAR(200) NOT NULL,
  lastname   VARCHAR(200) NOT NULL,
  email      VARCHAR(100) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO user
  (firstname, lastname, email)
VALUES
  ('Weerachai', 'Ruengrangsan', 'wee.ru@gmail.com'),
  ('Paweena', 'Suksawad', 'paw.suk@gmail.com'),
  ('Surawat ', 'Pongpanitch ', 'su.po@gmail.com');