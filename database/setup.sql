Cuser_idREATE TABLE users (
    id serial PRIMARY KEY,
    username text NOT NULL UNIQUE,
    first_name text NOT NULL,
    last_name text NOT NULL,
    password text NOT NULL
);

CREATE TABLE loginsessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id integer REFERENCES users(id) NOT NULL,
    expires_at timestamp NOT NULL
);

CREATE TABLE exercises (
    name text NOT NULL,
    user_id integer REFERENCES users(id) NOT NULL,
    PRIMARY Key (name, user_id)
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
    exercise text REFERENCES exercises(name, user_id) NOT NULL,
    comment text
);

CREATE TABLE sets (
    id serial PRIMARY KEY,
    lift_id integer REFERENCES lifts(id) NOT NULL,
    setnr integer NOT NULL,
    reps decimal NOT NULL,
    wieght decimal NOT NULL
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
