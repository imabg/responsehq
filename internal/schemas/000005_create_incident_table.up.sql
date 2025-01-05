CREATE TABLE IF NOT EXISTS incidents
(
    id            TEXT      NOT NULL PRIMARY KEY,
    name          TEXT      NOT NULL,
    is_backfilled bool DEFAULT FALSE,
    body          TEXT,
    is_active     BOOL DEFAULT TRUE,
    page_id       TEXT      NOT NULL,
    company_id    TEXT      NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (page_id) REFERENCES pages (id) ON DELETE CASCADE,
    created_at    TIMESTAMP NOT NULL,
    updated_at    TIMESTAMP
);