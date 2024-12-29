CREATE TYPE history AS ENUM (
    '7',
    '14',
    '30',
    '90',
    '180',
    '365'
    );

CREATE TABLE IF NOT EXISTS status_page
(
    id         TEXT PRIMARY KEY,
    url           TEXT      NOT NULL UNIQUE,
    is_active     BOOLEAN            DEFAULT TRUE,
    support_url   TEXT,
    logo_url      TEXT,
    timezone      TEXT,
    history_shows history   NOT NULL,
    company_id TEXT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    created_at    timestamp NOT NULL DEFAULT NOW(),
    updated_at    timestamp
);