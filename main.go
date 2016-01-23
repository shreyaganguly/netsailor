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
)
func main() {
   flag.Parse()
   args := flag.Args()
   protocol := "tcp"
   if *udpMode {
     protocol = "udp"
   }
   if *tlsMode {
     protocol = "tls"
   }
   if *listen {
     if len(args) == 0 || len(args) > 1{
       log.Fatal("Please provide port number along with optional flags")
     } else {
       Listener(args[0], protocol,*verbose)
     }
   } else if len(args) == 2 {
      Client(args[0],args[1],protocol,*verbose)
   } else if len(args) != 2 {
          log.Fatal("Please provide hostname and port number with optional flags")
     } else {
     flag.Usage()
   }

}
