DROP DATABASE IF EXISTS moviles;
CREATE DATABASE moviles;
USE moviles;

CREATE TABLE celulares(
id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
precio DECIMAL(16,2) NOT NULL,
descripcion VARCHAR(450) NOT NULL,
marca VARCHAR(200) NOT NULL,
modelo VARCHAR(200) NOT NULL,
lanzamiento DATE NOT NULL,
created_at TIMESTAMP NOT NULL
);