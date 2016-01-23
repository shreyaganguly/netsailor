package main

import (
  "log"
  "fmt"
  "strconv"
)

func GetHost(hostname string, port string) string {
  _, err := strconv.Atoi(port)
  if err != nil {
    log.Fatal("port number invalid: ",port)
  }
  return (fmt.Sprintf("%s:%s",hostname,port))
}

func SelectChannel(ch1, ch2 <-chan bool) {
  select {
  case <-ch1:
    log.Println("Connection closed from local process")
  case <-ch2:
    log.Println("Connection closed from remote process")
  }
}
