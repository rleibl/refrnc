
--
-- Docker (local, foreground, autoremove)
--     docker run -it --rm --name pg -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
--
-- Database connection
--     psql -U postgres -5432 -h localhost
--
CREATE USER test;
ALTER ROLE test WITH PASSWORD 'password';
CREATE DATABASE test WITH OWNER test;
