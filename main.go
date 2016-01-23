package main

import (
  "flag"
)

var (
  hostname  = flag.String("h","localhost","listen on host")
  port      = flag.String("p","8000","listen on port")
  listen    = flag.Bool("l",false,"to enable the listen mode")
  udpMode   = flag.Bool("u",false,"udp protocol")
  tlsMode   = flag.Bool("s", false, "tls protocol")
)
func main() {
   flag.Parse()
   protocol := "tcp"
   if *udpMode {
     protocol = "udp"
   }
   if *tlsMode {
     protocol = "tls"
   }
   if *listen {
      Listener(*port, protocol)
   } else if *hostname != "" {
      Client(*hostname,*port,protocol)
   } else {
     flag.Usage()
   }

}
