CREATE DATABASE IF NOT EXISTS db_user ENGINE = Atomic COMMENT 'User database';

CREATE TABLE IF NOT EXISTS db_user.users
(
    UserId UInt8 NOT NULL,
    Email String NOT NULL
) ENGINE = MergeTree()
ORDER BY UserId; 