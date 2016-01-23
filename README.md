# netsailor

to create certificate files (server.key and server.pem)

Generated private key

openssl genrsa -out server.key 2048
To generate a certificate

openssl req -new -x509 -key server.key -out server.pem -days 3650


## TODO
to allow sending of zip files or package files(not supported as of now)
