CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    email varchar UNIQUE,
    password_hash varchar,
    name varchar
);

CREATE TABLE IF NOT EXISTS items (
    id uuid PRIMARY KEY,
    item_description varchar NOT NULL,
    user_id uuid REFERENCES users (id) NOT NULL,
    done boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS oauth_apps (
    id uuid PRIMARY KEY,
    client_id varchar NOT NULL UNIQUE,
    client_secret varchar NOT NULL UNIQUE,
    name varchar NOT NULL,
    user_id uuid REFERENCES users (id) NOT NULL,
    redirect_uri varchar
);

CREATE TABLE IF NOT EXISTS oauth_authorizations (
    id uuid PRIMARY KEY,
    app uuid REFERENCES oauth_apps (id) NOT NULL,
    user_id uuid REFERENCES users (id) NOT NULL,
    code varchar NOT NULL,
    state varchar,
    scopes varchar[]
);

CREATE TYPE token_type AS ENUM ('oauth', 'client');
CREATE TYPE client_type AS ENUM ('mobile', 'web', 'personal', 'other');

CREATE TABLE IF NOT EXISTS tokens (
    id uuid PRIMARY KEY,
    token varchar NOT NULL UNIQUE,
    scopes varchar[],
    valid boolean DEFAULT TRUE,
    created_on timestamp DEFAULT NOW(),
    token_type token_type NOT NULL,
    user_id uuid REFERENCES users (id) NOT NULL,
    client_type client_type CHECK (
        CASE
            WHEN (token_type = 'client' AND client_type IS NOT NULL) THEN TRUE
            WHEN (token_type = 'oauth' AND client_type IS NULL) THEN TRUE
            ELSE FALSE
        END
    ),
    oauth_app uuid REFERENCES oauth_apps (id) CHECK (
        CASE
            WHEN (token_type = 'oauth' AND oauth_app IS NOT NULL) THEN TRUE
            WHEN (token_type = 'client' AND oauth_app IS NULL) THEN TRUE
            ELSE FALSE
        END
    )
);
