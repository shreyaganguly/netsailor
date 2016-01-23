## Net-sailor
The Net-Sailor is a simple implementation of the netcat utility in go that allows to listen and send data over SSL,TCP and UDP protocols.

### Examples:

**$ ns hostname 42**

Open a TCP connection to port 42 of hostname.

**$ ns -u hostname 53**

Open a UDP connection to port 53 of hostname.

**$ ns -v hostname 53**

Open a TCP connection to port 53 of hostname with connection logs.

**$ ns -s hostname 53**

Open a SSL connection to port 53 of hostname.

**$ ns -l 3000**

Listen on TCP port 3000, and once there is a connection, send stdin to the remote host, and send data from the remote host to stdout.





## TODO
to allow sending of zip files or package files(not supported as of now)
