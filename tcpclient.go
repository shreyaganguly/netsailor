package main

import (
  "net"
  "log"
)

func TCPClient(hostname string, port string)  {
  log.Println("Client running")
  host := GetHost(hostname, port)
  con, err := net.Dial("tcp", host)
  if err != nil {
    log.Println("ERROR: ",err)
    return
  }
  HandleTCPConnection(&con)

}
