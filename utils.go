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
