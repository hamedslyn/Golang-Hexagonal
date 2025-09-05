CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS todo_items (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    description text NOT NULL,
    due_date    timestamp NOT NULL
);
