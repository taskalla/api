CREATE TABLE IF NOT EXISTS users (
    id varchar PRIMARY KEY,
    email varchar UNIQUE,
    password_hash varchar,
    name varchar
);

CREATE TABLE IF NOT EXISTS items (
    id varchar PRIMARY KEY,
    title varchar,
    item_description varchar,
    user_id varchar REFERENCES users (id) NOT NULL
);

CREATE TYPE token_type AS ENUM ('oauth', 'client');

CREATE TABLE IF NOT EXISTS tokens (
    id varchar PRIMARY KEY,
    token varchar,
    scopes varchar[],
    valid boolean,
    created_on timestamp,
    token_type token_type,
    user_id varchar REFERENCES users (id) NOT NULL
);