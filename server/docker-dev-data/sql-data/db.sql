
-- users table
create table users (
  id serial,
  fullname varchar(80),
  email varchar(200),
  password varchar(70),
  createdAt timestamp DEFAULT CURRENT_TIMESTAMP,
  primary key(id)
);
