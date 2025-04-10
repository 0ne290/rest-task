CREATE TABLE tasks (
    uuid UUID PRIMARY KEY,
    user_uuid UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);