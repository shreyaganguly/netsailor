package main

import (
  "flag"
)

var (
  hostname = flag.String("h","localhost","listen on host")
  port     = flag.String("p","8000","listen on port")
  listen  = flag.Bool("l",false,"to enable the listen mode")
  proto   = flag.String("proto","tcp","check for which protocol to use")
)
func main() {
   flag.Parse()
   if *listen {
      TLSListener(*port)
   } else {
     TLSClient(*hostname,*port)
   }

}
