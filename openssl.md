# openssl

## server
```
    openssl s_server -key <privatekey> -cert <cert> -accept <port>
```
Note that keys are stored such that only root can see them, so bash
autocompletion may not work return html text with cypher information
```
    openssl s_server -key key.pem -cert <cert> -accept fullchain.pem -www
```

## client
```
   openssl s_client -connect forum.haufe.de:443
```

## Server Certificate
With SNI
```
    openssl s_client -showcerts -servername www.example.com -connect www.example.com:443 </dev/null
```
Without SNI
```
    openssl s_client -showcerts -connect www.example.com:443 </dev/null
```
Full Cert
```
    echo | openssl s_client -servername www.example.com -connect www.example.com:443 2>/dev/null | openssl x509 -text
```
