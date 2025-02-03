-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS balances (
    user_id UUID PRIMARY KEY,
    balance NUMERIC DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_user_id UUID,
    to_user_id UUID,
    amount NUMERIC NOT NULL,
    type TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL
);

INSERT INTO users (id, name) VALUES 
    ('73222d73-b663-4c71-9e2a-ec84d5040ed1','User 1'),
    ('73222d73-b663-4c71-9e2a-ec84d5040ed2','User 2');

INSERT INTO balances (user_id, balance) 
VALUES 
    ('73222d73-b663-4c71-9e2a-ec84d5040ed1', 100),
    ('73222d73-b663-4c71-9e2a-ec84d5040ed2', 200);

-- +goose Down
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS balances;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS pgcrypto;
