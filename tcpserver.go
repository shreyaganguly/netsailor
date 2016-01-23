package main

import (
  "net"
  "log"
)

func TCPListener(host string, port string) {
  log.Println("Listener Running")
  hostname := GetHost(host,port)
  listener, err := net.Listen("tcp", hostname)
  	if err != nil {
  		log.Println("ERROR :", err)
  	}
  	defer listener.Close()

  	conn, err := listener.Accept()
  	if err != nil {
  		log.Fatal(err)
  	}

  	HandleTCPConnection(conn)
}
