CREATE TABLE IF NOT EXISTS companies
(
    id         TEXT PRIMARY KEY,
    name            TEXT      NOT NULL,
    created_by TEXT NOT NULL,
    is_active       BOOLEAN            DEFAULT TRUE,
    subscription_id INT       NOT NULL,
    FOREIGN KEY (subscription_id) REFERENCES subscriptions (id),
    created_at      timestamp NOT NULL DEFAULT NOW(),
    updated_at      timestamp
);
