ALTER DATABASE postgres SET timezone TO 'Europe/Moscow';

CREATE TABLE users
(
    id          SERIAL       NOT NULL UNIQUE,
    age         INT NOT NULL,
    email       VARCHAR(50)  NOT NULL UNIQUE
);

COMMENT ON COLUMN users.age IS 'user age';
COMMENT ON COLUMN users.email IS 'user email';