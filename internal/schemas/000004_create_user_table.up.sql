CREATE TABLE IF NOT EXISTS users
(
    id         uuid PRIMARY KEY,
    email      TEXT      NOT NULL,
    company_id uuid      NOT NULL UNIQUE,
    name       TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    UNIQUE (email, company_id),
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp
);