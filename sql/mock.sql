-- This SQL script fills the database with mock data, and is automatically run on the first startup of a dev database

-- generate a mock user
INSERT INTO users (id, email, password_hash, name) VALUES (
    'd2c5c3ad-2e6a-45e2-941b-4427a00714f4',
    'taskalla@example.com',
    /* bcrypt hash of "password" */'$2y$12$E2/CCi8k3cceNNMbuMlzS.nUT8Tjjet2oIgsmMKmtWBFewup8arcq',
    'Test User'
);

-- generate some mock items
INSERT INTO items (id, item_description, user_id) VALUES (
    '4ca9b7f9-fe85-48e1-9765-d7d5cfcb9e3e',
    'Get eggs',
    'd2c5c3ad-2e6a-45e2-941b-4427a00714f4'
);

INSERT INTO items (id, item_description, user_id) VALUES (
    '8b06f951-b02f-4296-9aa9-f877fb3395a3',
    'Feed the dog',
    'd2c5c3ad-2e6a-45e2-941b-4427a00714f4'
);

INSERT INTO items (id, item_description, user_id, done) VALUES (
    'ae768069-8c09-4ab3-b6f1-69abdba033ec',
    'Learn to code',
    'd2c5c3ad-2e6a-45e2-941b-4427a00714f4',
    true
);
