## Net-sailor
The Net-Sailor is a simple implementation of the netcat utility in go that allows to listen and send data over SSL,TCP and UDP protocols.

### Motivation
The conventional netcat that is used doesn't support ssl connection or any kind of authentication as such.So I thought it will be of use to add SSL protocol and thereby it intrigued me to design netcat command utility purely in golang and adding extra functionality of supporting ssl protocols.The name net sailor itself comes from the fact that this command line utility is sails across tcp,udp and tls(ssl) protocols.

### Installation

- Just run :
  ```
  $ go build
  ```

### Usage:

### Client Usage:

```
netsailor [-u] [-s] [-a] [-b] [-z] [-v] [hostname] [port]
```
### Listener Usage:

```
netsailor [-l] [-u] [-s] [-c] [-n] [-v] [port]
```

### Description:

The utility allows to listen UDP\TCP\SSl ports and send data to remote ports over TCP\UDP\SSL. Main usage scenario is testing network protocols and accessibility of the open ports.

The options are as follows:

``` -l ```
	Used to specify that netsailor should listen for an incoming connection rather than initiate a connection to a remote host.

``` -u ```
	Use UDP instead of the default option of TCP.

``` -s ```
	Use SSL instead of the default option of TCP.

``` -v ```
	Have netsailor give more verbose output.

``` -a ```
	Allow authentication mode for client side(allowing client to verify the server by the servername).

``` -b  servername ```
	Have client mention the servername(has to be only used when -a mode is on) to verify the server.

``` -z  portrange ```
	Have client mention the range of ports to be scanned to check whether they are open or not(should be used with tcp and ssl mode only).

``` -c  path/to/cert/files ```
	Have listener mention the path to the certificate files(.key and .pem file) in SSL connection(to be used only when -s flag is set)(default is present working directory).

``` -n  certificatename ```
	Have listener mention the name of the certificate file(.key and .pem file) in SSL connection(to be used only when -s flag is set)(default name is server.key and server.pem).



### Examples:

## Client Side:

**$ netsailor hostname 42**

Open a TCP connection to port 42 of hostname.

**$ netsailor -u hostname 53**

Open a UDP connection to port 53 of hostname.

**$ netsailor -v hostname 53**

Open a TCP connection to port 53 of hostname in verbose with connection logs.

**$ netsailor -s hostname 53**

Open a SSL connection to port 53 of hostname.

**$ netsailor -s -a -b servername hostname 8000**

Open a SSL connection to port 8000 of hostname and verify servername before sending stdin to remote host, and send data from remote host to stdout.

**$ netsailor -z hostname 8000-8003**

It will scan all the open ports from 8000 to 8003 If -s is mentioned it will scan for all open SSL ports within the specified range else it will scan for open tcp ports.

## Listener Side:

**$ netsailor -l 3000**

Listen on TCP port 3000, and once there is a connection, send stdin to the remote host, and send data from the remote host to stdout.

**$ netsailor -u -l 3000**

Listen on UDP port 3000, and once there is a connection, send stdin to the remote host, and send data from the remote host to stdout.

**$ netsailor -s -l 3000**

Listen on SSL port 3000, and once there is a connection, send stdin to the remote host, and send data from the remote host to stdout.

**$ netsailor -l -v 3000**

Listen on TCP port 3000, and once there is a connection, send stdin to the remote host, and send data from the remote host to stdout in verbose mode with connection logs.

**$ netsailor -l -s -c path/to/cert/files -n certname 8000**

Listen on SSL port 8000, and thereby providing the Certificates from specified path(default present working directory) and specifying the name(default "server")


### File Transfer Examples:

Start by using nc to listen on a specific port, with output captured into a file:

**netsailor -l 1234 > filename.out**

Using a second machine, connect to the listening nc process, feeding it the file which is to be tranetsailorferred:

**netsailor host.example.com 1234 < filename.in**

After the file has been transferred, the connection will close automatically.
