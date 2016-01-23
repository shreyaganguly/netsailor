package main

import (
  "net"
  "log"
  "crypto/tls"
)

func Client(hostname string, port string, protocol string) {
  switch protocol {
  case "tcp" :TCPClient(hostname,port)
  case "tls" :TLSClient(hostname,port)
  case "udp" :UDPClient(hostname,port)
  }
}
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

func UDPClient(hostname string, port string) {
  log.Println("Client running")
  host := GetHost(hostname,port)
  connectAddr, errAddr := net.ResolveUDPAddr("udp", host)
  if errAddr != nil {
    log.Println("ERROR: ",errAddr)
    return
  }
  con, err := net.DialUDP("udp",nil,connectAddr)
  if err != nil {
    log.Println("ERROR: ",err)
    return
  }
  HandleUDPConnection(con)
}
