ALTER DATABASE postgres SET timezone TO 'Europe/Moscow';

CREATE TYPE gender AS ENUM ('жеребец', 'кобыла', 'мерин');

CREATE TYPE habit AS ENUM ('нервная', 'спокойная', 'добросовестная', 'жизнерадостная');

CREATE TABLE breeds
(
    id          SERIAL       NOT NULL UNIQUE,
    name        VARCHAR(50)  NOT NULL
);

CREATE TABLE horses
(
    id            SERIAL       NOT NULL UNIQUE,
    name          VARCHAR(20)  NOT NULL,
    gender        gender,
    age           INT NOT NULL,
    weight        INT NOT NULL,
    breed_id      INT REFERENCES breeds (id) ON DELETE CASCADE NOT NULL,
    birthplace    VARCHAR(100),
    habit         habit,
    rentalprice   INT NOT NULL,
    purchaseprice INT NOT NULL,
    sold          BOOLEAN      NOT NULL DEFAULT FALSE
);

COMMENT ON COLUMN horses.age IS 'years';
COMMENT ON COLUMN horses.weight IS 'kilograms';
COMMENT ON COLUMN horses.rentalprice IS 'rubles';
COMMENT ON COLUMN horses.purchaseprice IS 'rubles';

CREATE TABLE tasks
(
    id          SERIAL       NOT NULL UNIQUE,
    title       VARCHAR(50)  NOT NULL
);

CREATE TABLE horses_tasks
(
    horse_id    INT REFERENCES horses (id)      NOT NULL,
    task_id   INT REFERENCES tasks (id)     NOT NULL,
    UNIQUE (horse_id, task_id)
);

CREATE TABLE rentals
(
    id          SERIAL                          NOT NULL UNIQUE,
    horse_id    INT REFERENCES horses (id)      NOT NULL,
    since       TIMESTAMP,
    till        TIMESTAMP,
    paid        BOOLEAN                         NOT NULL DEFAULT FALSE
);