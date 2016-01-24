package main

import (
  "flag"
  "log"
)

var (
  listen    = flag.Bool("l",false,"to enable the listen mode")
  udpMode   = flag.Bool("u",false,"udp protocol")
  tlsMode   = flag.Bool("s", false, "tls protocol")
  verbose   = flag.Bool("v", false, "turn on verbose mode for descriptive logs")
  authorize = flag.Bool("a",false,"verify server side(should not be used by the listener server)")
	servername = flag.String("b","","mention the server name(must be used if -a flag is on the client side)")
  certlocation = flag.String("c",getWorkingDirectory(),"mention the path where .key and .pem files are located(to be only used by the listener when -s mode is on)(default is present working directory)")
  certname     = flag.String("n", "server", "mention the name of the pem file and key file(flag should be the same name)(for the listener only)(default name server.key and server.pem)")
  scan         = flag.Bool("z", false, "just scan for listening daemons(valid for tls and tcp protocols only), without sending any data to them(should not be used with -l flag)")
)
func main() {
   flag.Parse()
   args := flag.Args()
   protocol := "tcp"
   if *udpMode && *tlsMode {
     log.Fatal("UDP and TLS cannot be used together to specify the protocol")
   }
   if *udpMode {
     protocol = "udp"
   }
   if *tlsMode {
     protocol = "tls"
   }
   if *listen {
     if (*authorize || *servername != "") || (!(*tlsMode) && *certlocation != getWorkingDirectory()) || *scan {
       flag.Usage()
       return
     }
     if len(args) == 0 || len(args) > 1 {
       log.Fatal("Please provide port number along with optional flags")
     } else {
       Listener(args[0], protocol,*verbose)
     }
   } else if (*authorize && *servername == "") {
       flag.Usage()
       return
    }  else if len(args) == 2 {
      Client(args[0],args[1],protocol,*verbose)
   } else if len(args) != 2 {
          log.Fatal("Please provide hostname and port number with optional flags")
     } else {
     flag.Usage()
     return
   }

}
