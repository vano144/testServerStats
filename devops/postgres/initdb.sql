CREATE ROLE user_stats
LOGIN
PASSWORD 'postgres'
NOSUPERUSER
NOINHERIT
NOCREATEDB
NOCREATEROLE
NOREPLICATION;

CREATE DATABASE user_stats
    WITH OWNER = postgres
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

ALTER ROLE user_stats IN DATABASE user_stats SET search_path = user_stats;


\c user_stats
CREATE SCHEMA user_stats;

ALTER SCHEMA user_stats
    OWNER TO user_stats;