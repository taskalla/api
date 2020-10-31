CREATE TABLE IF NOT EXISTS users (
    id varchar PRIMARY KEY,
    email varchar UNIQUE,
    password_hash varchar,
    name varchar
);

CREATE TABLE IF NOT EXISTS items (
    id varchar PRIMARY KEY,
    title varchar,
    item_description varchar NOT NULL,
    user_id varchar REFERENCES users (id) NOT NULL
);

CREATE TYPE token_type AS ENUM ('oauth', 'client');
CREATE TYPE client_type AS ENUM ('mobile', 'web', 'personal');

CREATE TABLE IF NOT EXISTS tokens (
    id varchar PRIMARY KEY,
    token varchar NOT NULL UNIQUE,
    scopes varchar[],
    valid boolean DEFAULT TRUE,
    created_on timestamp,
    token_type token_type NOT NULL,
    user_id varchar REFERENCES users (id) NOT NULL,
    client_type client_type CHECK (
        CASE
            WHEN (token_type = 'client' AND client_type IS NOT NULL) THEN TRUE
            WHEN (token_type = 'oauth' AND client_type IS NULL) THEN TRUE
            ELSE FALSE
        END
    )
);