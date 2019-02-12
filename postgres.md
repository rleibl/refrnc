# postgres

## new database
```
    PSQL="psql -h localhost -p 5432 -U postgres"
    $PSQL -c "CREATE USER dbuser"
    $PSQL -c "ALTER ROLE dbuser WITH PASSWORD 'dbusers_password'"
    $PSQL -c "CREATE DATABASE dbtest WITH OWNER dbuser"
```

## docker (see also [docker](./docker.md))
```
    docker run --name pg -e POSTGRES_PASSWORD=pg_password -p 5432:5432 -d postgres
```
if you need foreground, remember "-it --rm"
