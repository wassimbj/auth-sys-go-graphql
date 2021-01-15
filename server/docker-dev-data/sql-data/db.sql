
-- users table
create table users (
  id serial,
  fullname varchar(80),
  email varchar(200),
  password varchar(50),
  image varchar(250),
  createdAt timestamp DEFAULT CURRENT_TIMESTAMP,
  primary key(id)
);

-- posts table
create table posts (
  id serial,
  slug varchar(250),
  title varchar(250),
  body varchar(5000),
  author integer,
  tags text[], -- array of string
  createdAt timestamp DEFAULT CURRENT_TIMESTAMP,
  category varchar(100),
  primary key(id)
);
