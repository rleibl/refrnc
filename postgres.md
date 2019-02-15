# Postgres - Reference
## new database
```
    PSQL="psql -h localhost -p 5432 -U postgres"
    $PSQL -c "CREATE USER dbuser"
    $PSQL -c "ALTER ROLE dbuser WITH PASSWORD 'dbusers_password'"
    $PSQL -c "CREATE DATABASE dbtest WITH OWNER dbuser"
```

## list databases
```
psql -l
postgres=# \l
```

## create database
```
sudo -u postgres createdb foo
sudo su postgres; createdb foo
```

## list tables
```
# \dt
```

## user / priviledge management
```
sudo -u postgres createuser <role>
sudo -u postgres psql <database>
db=# grant select, insert, update, delete on <table> to <role>
db=# alter role <role> with password '<password>'
```

## connect to local DB
```
psql -h localhost -U <user> <database>

psql -h localhost -p 5432 -U postgres
```

## database and user with psql
```
PSQL="psql -h localhost -p 5432 -U postgres"

# create user
$PSQL -c "CREATE USER dbuser"
$PSQL -c "ALTER ROLE dbuser WITH PASSWORD 'pppp'"

# create database
$PSQL -c "CREATE DATABASE dbtest WITH OWNER dbuser"
```


## docker (see also [docker](./docker.md))
```
    docker run --name pg -e POSTGRES_PASSWORD=pg_password -p 5432:5432 -d postgres
```
if you need foreground, remember "-it --rm"

