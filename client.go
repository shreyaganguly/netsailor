package main

import (
  "net"
  "log"
  "crypto/tls"
)

func Client(hostname string, port string, protocol string,verbose bool) {
  if verbose {
		log.Println("Dialing", protocol, "on", hostname, ":", port)
	}
  switch protocol {
  case "tcp" :TCPClient(hostname,port,verbose)
  case "tls" :TLSClient(hostname,port,verbose)
  case "udp" :UDPClient(hostname,port,verbose)
  }
}

func TCPClient(hostname string, port string,verbose bool)  {
  if *scan {
    checkForOpenPorts(hostname, port, "tcp", verbose)
    return
  }
  host := GetHost(hostname, port)
  con, err := net.Dial("tcp", host)
  if err != nil {
    if verbose {
        log.Println("ERROR: ",err)
    }
    return
  }
  HandleTCPConnection(&con,verbose)
}

func TLSClient(hostname string, port string,verbose bool) {
  if *scan {
    checkForOpenPorts(hostname, port , "tls", verbose)
    return
  }
  host := GetHost(hostname,port)
  var conf *tls.Config
		if *authorize {
		conf = &tls.Config{
			ServerName: *servername,
		}
   } else {
			conf = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
  con, err := tls.Dial("tcp",host,conf)
  if err != nil {
    if verbose {
        log.Println("ERROR: ",err)
    }
    return
  }
  HandleTLSConnection(con,verbose)
}

func UDPClient(hostname string, port string,verbose bool) {
  if *scan {
    log.Fatal("OOPS!! this is not implemented yet")
  }
  host := GetHost(hostname,port)
  connectAddr, errAddr := net.ResolveUDPAddr("udp", host)
  if errAddr != nil {
    if verbose {
        log.Println("ERROR: ",errAddr)
    }
    return
  }
  con, err := net.DialUDP("udp",nil,connectAddr)
  if err != nil {
    if verbose {
        log.Println("ERROR: ",err)
    }
    return
  }
  HandleUDPConnection(con,verbose)
}
