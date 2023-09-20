-- Add up migration script here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    "id"         uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name"       varchar          NOT NULL,
    "image"      varchar          NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE  DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE  DEFAULT NOW()
)