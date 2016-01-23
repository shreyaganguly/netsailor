package main

import (
  "log"
  "net"
  "os"
  "bufio"
  "io"
)

func HandleTCPConnection(con net.Conn) {
  chan_local := readAndWrite(os.Stdin, con)
  chan_remote := readAndWrite(con, os.Stdout)
  select {
    case <- chan_local:
      log.Println("Connection closed from local process")
    case <- chan_remote:
      log.Println("Connection closed from remote process")
  }
}

func readAndWrite(r io.Reader, w io.Writer) <-chan bool  {
  c := make(chan bool)
  go func() {
    defer func() {
      if con, ok := w.(net.Conn); ok {
          con.Close()
      }
      c <- false
    }()
    for {
      message, errRead := bufio.NewReader(r).ReadString('\n')
      if errRead != nil {
        if errRead != io.EOF {
          log.Println("READ ERROR: ",errRead)
        }
        break
      }
      writer := bufio.NewWriter(w)
      _, errWrite := writer.WriteString(message)
      writer.Flush()
      if errWrite != nil {
        log.Println("WRITE ERROR: ",errWrite)
        return
      }
    }
  }()
  return c
}
