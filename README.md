# Processando linhas do MySQL com Go Routines

```sh
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=senha -d mysql:latest
```

```sql
create database projeto;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255)
);

INSERT INTO users (name, email) VALUES ('Ricardo','ricardo@gmail.com');
INSERT INTO users (name, email) VALUES ('João','joao@gmail.com');
INSERT INTO users (name, email) VALUES ('Maria','maria@gmail.com');
INSERT INTO users (name, email) VALUES ('Pedro','pedro@gmail.com');
INSERT INTO users (name, email) VALUES ('Ana','ana@gmail.com');
INSERT INTO users (name, email) VALUES ('Fernanda','fernanda@gmail.com');
INSERT INTO users (name, email) VALUES ('Lucas','lucas@gmail.com');
INSERT INTO users (name, email) VALUES ('Gabriel','gabriel@gmail.com');
INSERT INTO users (name, email) VALUES ('Carla','carla@gmail.com');
INSERT INTO users (name, email) VALUES ('Renato','renato@gmail.com');
INSERT INTO users (name, email) VALUES ('Julia','julia@gmail.com');
INSERT INTO users (name, email) VALUES ('Thiago','thiago@gmail.com');
INSERT INTO users (name, email) VALUES ('Isabella','isabella@gmail.com');
INSERT INTO users (name, email) VALUES ('Marcos','marcos@gmail.com');
INSERT INTO users (name, email) VALUES ('Laura','laura@gmail.com');
INSERT INTO users (name, email) VALUES ('Matheus','matheus@gmail.com');
INSERT INTO users (name, email) VALUES ('Bruna','bruna@gmail.com');
INSERT INTO users (name, email) VALUES ('Rafael','rafael@gmail.com');
INSERT INTO users (name, email) VALUES ('Camila','camila@gmail.com');
INSERT INTO users (name, email) VALUES ('Vinicius','vinicius@gmail.com');
INSERT INTO users (name, email) VALUES ('Amanda','amanda@gmail.com');
INSERT INTO users (name, email) VALUES ('Gustavo','gustavo@gmail.com');
INSERT INTO users (name, email) VALUES ('Larissa','larissa@gmail.com');
INSERT INTO users (name, email) VALUES ('Rodrigo','rodrigo@gmail.com');
INSERT INTO users (name, email) VALUES ('Paola','paola@gmail.com');
INSERT INTO users (name, email) VALUES ('Felipe','felipe@gmail.com');
INSERT INTO users (name, email) VALUES ('Carolina','carolina@gmail.com');
INSERT INTO users (name, email) VALUES ('Henrique','henrique@gmail.com');
INSERT INTO users (name, email) VALUES ('Vanessa','vanessa@gmail.com');
INSERT INTO users (name, email) VALUES ('Diego','diego@gmail.com');

SELECT id, name, email FROM users;

```