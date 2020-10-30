CREATE TABLE IF NOT EXISTS items (
    title varchar,
    description varchar
);

CREATE TABLE IF NOT EXISTS users (
    email varchar,
    password_hash varchar,
    id integer PRIMARY KEY,
    name varchar
);

CREATE TYPE token_type AS ENUM ('oauth', 'client');

CREATE TABLE IF NOT EXISTS tokens (
    scopes varchar[],
    token varchar,
    valid boolean,
    created_on timestamp,
    token_type token_type,
    user varchar
);