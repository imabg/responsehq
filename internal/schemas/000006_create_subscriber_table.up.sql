CREATE TYPE subscriber_type AS ENUM (
    'mail',
    'webhook',
    'slack',
    'ms_teams'
    );

CREATE TABLE IF NOT EXISTS subscribers
(
    id          TEXT            NOT NULL PRIMARY KEY,
    type        subscriber_type NOT NULL,
    value       TEXT            NOT NULL,
    is_verified BOOL                     DEFAULT FALSE,
    is_archived BOOL                     DEFAULT FALSE,
    page_id     TEXT            NOT NULL,
    FOREIGN KEY (page_id) REFERENCES pages (id) ON DELETE CASCADE,
    created_at  TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP
);