# docker

## options
* -i - interactive
* -t - tty interface
* --rm - remove container when stopped
* -p dockerport:remoteport - export dockerport from the container to remoteport on the host

## docker-machine
```
        docker-machine --tls-ca-cert=ca.pem --tls-client-cert=cert.pem --tls-client-key=key.pem create --driver none --url tcp://docker-machine-host.com:2376 bla
        docker env bla
        eval $(docker env bla)
```

## postgres
```
        docker run --name pg -e POSTGRES_PASSWORD=pppp -p 5432:5432 -d postgres
        psql -h localhost -p 5432 -U postgres
```

## mongodb
```
        docker run -it --rm -p 27017:27017 --name mymongo mongo
```
