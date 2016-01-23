package main

import (
  "net"
  "log"
  "crypto/tls"
)

func TCPListener(port string) {
  // log.Println("Listener Running")
  host := GetHost("",port)
  listener, err := net.Listen("tcp", host)
  	if err != nil {
  		log.Println("ERROR :", err)
  	}
  	defer listener.Close()

  	conn, err := listener.Accept()
  	if err != nil {
  		log.Fatal(err)
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
  }
  HandleTCPConnection(&conn)
}
