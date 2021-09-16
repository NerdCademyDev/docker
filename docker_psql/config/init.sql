create table post(
    id serial unique not null,
    title varchar(128),
    content text,
    primary key(id)
);

insert into post(title, content) values(
    'Hello Docker Compose with PostgreSQL', 
    'This is my first Post using Docker Compose with PostgreSQL'
);