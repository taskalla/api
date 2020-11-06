-- This SQL script fills the database with mock data, and is automatically run on the first startup of a dev database

-- generate a mock user
INSERT INTO users (id, email, password_hash, name) VALUES (
    'd2c5c3ad-2e6a-45e2-941b-4427a00714f4',
    'taskalla@example.com',
    /* bcrypt hash of "password" */'$2y$12$E2/CCi8k3cceNNMbuMlzS.nUT8Tjjet2oIgsmMKmtWBFewup8arcq',
    'Test User'
);

INSERT INTO items (id, title, item_description, user_id) VALUES (
    'hi',
    null,
    '',
    'd2c5c3ad-2e6a-45e2-941b-4427a00714f4'
);