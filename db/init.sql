CREATE DATABASE wiki;

CREATE TABLE [page] (
    id          SERIAL PRIMARY KEY,
    slug        varchar(40) NOT NULL,
    title       text,
    content     integer NOT NULL
);