CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    login TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_users_login ON users(login);
