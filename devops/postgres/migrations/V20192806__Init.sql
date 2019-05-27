CREATE TABLE users (
    id integer PRIMARY KEY,
    age integer,
    sex varchar(1)
);

CREATE TABLE stats_values (
    id SERIAL PRIMARY KEY,
    user_id integer,
    action varchar(50),
    ts TIMESTAMP WITH TIME ZONE NOT NULL
);