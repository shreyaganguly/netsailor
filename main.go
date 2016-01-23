package main

import (
  "log"
  "flag"
)

var (
  hostname = flag.String("b","localhost","listen on host")
  port     = flag.String("p","8000","listen on port")
  listen  = flag.Bool("l",false,"to enable the listen mode")
)
func main() {
   flag.Parse()
   if *listen {
      TCPListener(*hostname,*port)
   } else {
     TCPClient(*hostname,*port)
     log.Println("Client mode")
   }

}
