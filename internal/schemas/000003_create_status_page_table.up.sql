CREATE TYPE history AS ENUM (
    '7',
    '14',
    '30',
    '90',
    '180',
    '365'
    );

CREATE TABLE IF NOT EXISTS pages
(
    id         TEXT PRIMARY KEY,
    url           TEXT      NOT NULL UNIQUE,
    is_active     BOOLEAN            DEFAULT TRUE,
    support_url   TEXT,
    logo_url      TEXT,
    timezone      TEXT,
    history_shows history   NOT NULL,
    send_notification BOOLEAN NOT NULL,
    company_id TEXT NOT NULL,
    subscription_id   INT     NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (subscription_id) REFERENCES subscriptions (id),
    created_at    timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);