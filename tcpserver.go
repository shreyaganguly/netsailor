package main

import (
  "net"
  "log"
  "bufio"
)

func TCPListener(host string, port string) {
  log.Println("Listener Running")
  hostname := getHost(host,port)
  listener, err := net.Listen("tcp",hostname)
  if err != nil {
    log.Println("ERROR: ", err)
    return
  }
  defer listener.Close()
  log.Println("Listening to tcp at",hostname)
  for {
    con, err := listener.Accept()
    if err != nil {
      log.Println("ERROR: ", err)
      return
    }
    go func(c net.Conn) {
      for {
        message, _ := bufio.NewReader(c).ReadString('\n')
        log.Println(message)
      }
      }(con)
  }
}
