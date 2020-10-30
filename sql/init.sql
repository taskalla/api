CREATE TABLE IF NOT EXISTS items (
    title varchar,
    item_description varchar,
    parent_user varchar,
    id varchar PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users (
    email varchar,
    password_hash varchar,
    id varchar PRIMARY KEY,
    name varchar
);

CREATE TYPE token_type AS ENUM ('oauth', 'client');

CREATE TABLE IF NOT EXISTS tokens (
    scopes varchar[],
    token varchar,
    valid boolean,
    created_on timestamp,
    token_type token_type,
    parent_user varchar
);