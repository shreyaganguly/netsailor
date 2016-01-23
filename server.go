package main

import (
  "net"
  "log"
  "crypto/tls"
)

func Listener(port string,protocol string) {
  switch protocol {
  case "tcp": TCPListener(port)
  case "udp": UDPListener(port)
  case "tls": TLSListener(port)
  }
}
func TCPListener(port string) {
  log.Println("Listener Running")
  host := GetHost("",port)
  listener, err := net.Listen("tcp", host)
  	if err != nil {
  		log.Println("ERROR :", err)
      return
  	}
  	defer listener.Close()

  	conn, err := listener.Accept()
  	if err != nil {
      log.Println("ERROR :", err)
      return
  	}

  	HandleTCPConnection(&conn)
}

func TLSListener(port string) {
  log.Println("Listener running")
  host := GetHost("",port)
  cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
  if err != nil {
    log.Println("ERROR: ",err)
    return
  }
  config := &tls.Config{Certificates: []tls.Certificate{cer}}
  listener, err := tls.Listen("tcp",host,config)
  if err != nil {
    log.Println("ERROR: ", err)
    return
  }
  defer listener.Close()
  conn, err := listener.Accept()
  if err != nil {
    log.Println("ERROR: ",err)
    return
  }
  HandleTCPConnection(&conn)
}

func UDPListener(port string) {
  log.Println("Listener running")
  host := GetHost("",port)
  ServerAddr, errAddr := net.ResolveUDPAddr("udp", host)
  if errAddr != nil {
    log.Println("ERROR: ",errAddr)
    return
  }
  conn, err := net.ListenUDP("udp", ServerAddr)
  defer conn.Close()
  if err != nil {
    log.Println("ERROR: ",err)
    return
  }
  HandleUDPConnection(conn)
}
