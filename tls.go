package main

import (
 "bufio"
 "log"
 "io"
 "crypto/tls"
 "os"
)

func HandleTLSConnection(con *tls.Conn) {
  chan_local := readAndWriteTLS(bufio.NewReader(os.Stdin), bufio.NewWriter(con), con)
  chan_remote := readAndWriteTLS(bufio.NewReader(con), bufio.NewWriter(os.Stdout), con)
  SelectChannel(chan_local,chan_remote)
}

func readAndWriteTLS(r *bufio.Reader, w *bufio.Writer, con *tls.Conn) <-chan bool {
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
