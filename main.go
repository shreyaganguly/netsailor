package main

import (
  "log"
  "flag"
)

var (
  hostname = flag.String("b","localhost","listen on host")
  port     = flag.String("p","8000","listen on port")
)
func main() {
   flag.Parse()
   log.Println("HOST NAME is",*hostname)
   log.Println("PORT is",*port)
   TCPListener(*hostname,*port)
}
