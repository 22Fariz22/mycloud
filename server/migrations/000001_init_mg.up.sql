DROP TABLE IF EXISTS users CASCADE;
-- DROP TABLE IF EXISTS files CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE users
(
    user_id    UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    login      VARCHAR(32)              NOT NULL CHECK ( login <> '' ),
    password   VARCHAR(250)             NOT NULL CHECK ( octet_length(password) <> 0 )
    -- created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()

);

-- CREATE TABLE files
-- (
--     file_id    UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
--     -- user_id    VARCHAR,
--     title      VARCHAR(32)              NOT NULL,
--     data       bytea                    NOT NULL
--     -- created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
--     -- CONSTRAINT fk_user 
--     -- FOREIGN KEY(user_id)
--     -- REFERENCES users(user_id)
-- );

