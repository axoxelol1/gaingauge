CREATE TABLE users (
    id serial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL
);

CREATE TABLE exercises (
    name text PRIMARY KEY
);

CREATE TABLE workouts (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    date timestamp NOT NULL,
    comment text
);

CREATE TABLE lifts (
    id serial PRIMARY KEY,
    workout integer REFERENCES workouts(id) NOT NULL,
    exercise text REFERENCES exercises(name) NOT NULL,
    comment text
);

CREATE TABLE weighins (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    date timestamp NOT NULL,
    comment text
);

CREATE TABLE sleep_logs (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) NOT NULL,
    date timestamp NOT NULL,
    hours real NOT NULL,
    comment text
);
