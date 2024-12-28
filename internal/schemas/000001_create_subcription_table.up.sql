CREATE TYPE plans AS ENUM (
    'free',
    'basic',
    'enterprise'
    );

CREATE TABLE IF NOT EXISTS subscriptions
(
    id         BIGSERIAL PRIMARY KEY,
    is_active  BOOLEAN            DEFAULT TRUE,
    plan       plans     NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);
