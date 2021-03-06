CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name     VARCHAR(50) NOT NULL,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (256) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL
);