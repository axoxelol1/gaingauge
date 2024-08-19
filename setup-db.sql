CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username text NOT NULL UNIQUE,
    first_name text NOT NULL,
    last_name text NOT NULL,
    password text NOT NULL
);

CREATE TABLE loginsessions (
    auth string PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    expires_at timestamp NOT NULL
);

CREATE TABLE exercises (
    name text NOT NULL,
    user_id integer REFERENCES users(id) NOT NULL,
    PRIMARY KEY (name, user_id)
);

CREATE TABLE workouts (
    id INTEGER PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    date timestamp NOT NULL,
    comment text
);

CREATE TABLE lifts (
    id INTEGER PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    liftnr integer NOT NULL,
    workout integer REFERENCES workouts(id) NOT NULL,
    exercise text NOT NULL,
    comment text,
    FOREIGN KEY (exercise, user_id) REFERENCES exercises(name, user_id)
);

CREATE TABLE sets (
    id INTEGER PRIMARY KEY,
    lift integer REFERENCES lifts(id) NOT NULL,
    setnr integer NOT NULL,
    reps decimal NOT NULL,
    weight decimal NOT NULL
);

CREATE TABLE weighins (
    id INTEGER PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    date timestamp NOT NULL,
    comment text
);

CREATE TABLE sleep_logs (
    id INTEGER PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    date timestamp NOT NULL,
    hours real NOT NULL,
    comment text
);
