-- Add up migration script here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    "user_id"    uuid PRIMARY KEY    NOT NULL DEFAULT (uuid_generate_v4()),
    "name"       VARCHAR(255)        NOT NULL,
    "email"      VARCHAR(255) UNIQUE NOT NULL,
    "password"   VARCHAR(255)        NOT NULL,
    "image"      VARCHAR(255)        NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE     DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE     DEFAULT NOW()
);
-- Create a table for courses
CREATE TABLE IF NOT EXISTS courses
(
    course_id          uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    course_name        VARCHAR(255)     NOT NULL,
    course_description TEXT,
    video_url        VARCHAR(255),
    instructor_id      uuid REFERENCES users (user_id)


-- Foreign key referencing the instructor (user)
    -- Add other course-related fields as needed
);

-- Create a table to track which users are enrolled in which courses
CREATE TABLE IF NOT EXISTS user_courses
(
    enrollment_id uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    user_id       uuid REFERENCES users (user_id),    -- Foreign key referencing the user
    course_id     uuid REFERENCES courses (course_id) -- Foreign key referencing the course
);