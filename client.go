package main

import (
  "net"
  "log"
  "crypto/tls"
)

func TCPClient(hostname string, port string)  {
  //log.Println("Client running")
  host := GetHost(hostname, port)
  con, err := net.Dial("tcp", host)
  if err != nil {
    log.Println("ERROR: ",err)
    return
  }
  HandleTCPConnection(&con)
}

func TLSClient(hostname string, port string) {
  log.Println("Client running")
  host := GetHost(hostname,port)
  conf := &tls.Config{
    InsecureSkipVerify: true,
  }
  con, err := tls.Dial("tcp",host,conf)
  if err != nil {
    log.Println("ERROR: ",err)
    return 
  }
  HandleTLSConnection(con)
}
