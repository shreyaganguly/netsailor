package main

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"
	"strings"
)

func listenToPorts(host string, lower int, upper int, protocol string, verbose bool) {
	var err error
	for i := lower; i <= upper; i++ {
		hostname := GetHost(host, strconv.Itoa(i))
		switch protocol {
		case "tcp":
			_, err = net.Dial(protocol, hostname)
		case "tls":
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
			_, err = tls.Dial("tcp", hostname, conf)
		}
		if err != nil {
			if verbose {
				log.Println(err)
			}
		} else {
			log.Println("Connection succesful at", i, "with protocol", protocol)
		}
	}
}

func checkForOpenPorts(host string, portrange string, protocol string, verbose bool) {
	var lower, upper int
	var e error
	ports := strings.Split(portrange, "-")
	if len(ports) > 2 {
		log.Fatal("Please provide the correct range")
	}
	if len(ports) == 1 {
		lower, e = strconv.Atoi(ports[0])
		if e != nil {
			log.Fatal("Incorrect port number")
		}
		upper = lower
	} else {
		lower, e = strconv.Atoi(ports[0])
		if e != nil {
			log.Fatal("Incorrect port number")
		}
		upper, e = strconv.Atoi(ports[1])
		if e != nil {
			log.Fatal("Incorrect port number")
		}
	}
	listenToPorts(host, lower, upper, protocol, verbose)
}
