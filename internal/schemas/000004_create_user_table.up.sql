CREATE TABLE IF NOT EXISTS users
(
    id         TEXT PRIMARY KEY,
    email      TEXT      NOT NULL,
    company_id TEXT NOT NULL,
    sub_id     INT  NOT NULL,
    name       TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    FOREIGN KEY (sub_id) REFERENCES subscriptions (id),
    UNIQUE (email, company_id),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);