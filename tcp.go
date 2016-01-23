package main

import (
  "log"
  "net"
  "os"
  "bufio"
  "io"
)

func HandleTCPConnection(con *net.Conn) {
  chan_local := readAndWrite(bufio.NewReader(os.Stdin), bufio.NewWriter(*con), con)
  chan_remote := readAndWrite(bufio.NewReader(*con), bufio.NewWriter(os.Stdout), con)
  select {
    case <- chan_local:
      log.Println("Connection closed from local process")
    case <- chan_remote:
      log.Println("Connection closed from remote process")
  }
}

func readAndWrite(r *bufio.Reader, w *bufio.Writer, con *net.Conn) <-chan bool  {
  c := make(chan bool)
  go func() {
    defer func() {
          (*con).Close()
      c <- false
    }()
    for {
      message, errRead := r.ReadString('\n')
      if errRead != nil {
        if errRead != io.EOF {
          log.Println("READ ERROR: ",errRead)
        }
        break
      }
      _, errWrite := w.WriteString(message)
      w.Flush()
      if errWrite != nil {
        log.Println("WRITE ERROR: ",errWrite)
        return
      }
    }
  }()
  return c
}
