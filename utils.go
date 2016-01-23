package main

import (
  "log"
  "fmt"
  "strconv"
  "crypto/tls"
  "os"
)

func GetHost(hostname string, port string) string {
  _, err := strconv.Atoi(port)
  if err != nil {
    log.Fatal("port number invalid: ",port)
  }
  return (fmt.Sprintf("%s:%s",hostname,port))
}

func SelectChannel(ch1, ch2 <-chan bool,verbose bool) {
  select {
  case <-ch1:
    if verbose {
      log.Println("Connection closed from local process")
    }
  case <-ch2:
    if verbose {
      log.Println("Connection closed from remote process")
    }
  }
}

func getWorkingDirectory() string{
	wd,_ := os.Getwd()
	return wd
}
func getCertificate(certLocation string,verbose bool) *tls.Config {
  defer func() {
    _ = os.Chdir(getWorkingDirectory())
  }()
  cherr := os.Chdir(certLocation)
  if cherr != nil {
    log.Fatal("No such path exists")
  }
  if _, errkey := os.Stat("server.key"); os.IsNotExist(errkey) {
   log.Fatal("server.key file missing")
}
if _, errpem := os.Stat("server.pem"); os.IsNotExist(errpem) {
 log.Fatal("server.pem file missing")
}
  cer, err := tls.LoadX509KeyPair("server.pem", "server.key")
  if err != nil {
    if verbose {
      log.Println(err)
    }
    return nil
  }
  return (&tls.Config{Certificates: []tls.Certificate{cer}})
}
