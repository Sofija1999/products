CREATE TABLE users (
    id bigserial NOT NULL,
    first_name varchar NULL,
    last_name varchar NULL,
    email varchar NULL,
    password varchar NULL,
    created_at timestamp NULL,
    CONSTRAINT users_pk PRIMARY KEY(id)
);