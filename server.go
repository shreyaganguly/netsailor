package main

import (
  "net"
  "log"
  "crypto/tls"
)

func Listener(port string,protocol string,verbose bool) {
  if verbose {
    log.Println("Listening on", protocol, ":", port)
  }
  switch protocol {
  case "tcp": TCPListener(port,verbose)
  case "udp": UDPListener(port,verbose)
  case "tls": TLSListener(port,verbose)
  }
}
func TCPListener(port string,verbose bool) {
  host := GetHost("",port)
  listener, err := net.Listen("tcp", host)
  	if err != nil {
      if verbose {
          log.Println("ERROR: ",err)
      }
      return
  	}
  	defer listener.Close()

  	conn, err := listener.Accept()
  	if err != nil {
      if verbose {
          log.Println("ERROR: ",err)
      }
      return
  	}

  	HandleTCPConnection(&conn,verbose)
}

func TLSListener(port string,verbose bool) {
  host := GetHost("",port)
  config := getCertificate(*certlocation,verbose)
  listener, err := tls.Listen("tcp",host,config)
  if err != nil {
    if verbose {
        log.Println("ERROR: ",err)
    }
    return
  }
  defer listener.Close()
  conn, err := listener.Accept()
  if err != nil {
    if verbose {
        log.Println("ERROR: ",err)
    }
    return
  }
  HandleTCPConnection(&conn,verbose)
}

func UDPListener(port string,verbose bool) {
  host := GetHost("",port)
  ServerAddr, errAddr := net.ResolveUDPAddr("udp", host)
  if errAddr != nil {
    if verbose {
        log.Println("ERROR: ",errAddr)
    }
    return
  }
  conn, err := net.ListenUDP("udp", ServerAddr)
  if err != nil {
    if verbose {
        log.Println("ERROR: ",err)
    }
    return
  }
  defer conn.Close()
  HandleUDPConnection(conn,verbose)
}
