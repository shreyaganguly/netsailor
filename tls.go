package main

import (
 "bufio"
 "log"
 "io"
 "crypto/tls"
 "os"
)

func HandleTLSConnection(con *tls.Conn,verbose bool) {
  chan_local := readAndWriteTLS(bufio.NewReader(os.Stdin), bufio.NewWriter(con), con, verbose)
  chan_remote := readAndWriteTLS(bufio.NewReader(con), bufio.NewWriter(os.Stdout), con, verbose)
  SelectChannel(chan_local,chan_remote,verbose)
}

func readAndWriteTLS(r *bufio.Reader, w *bufio.Writer, con *tls.Conn,verbose bool) <-chan bool {
  c := make(chan bool)
  go func() {
    defer func() {
          (*con).Close()
      c <- false
    }()
    for {
      buf := make([]byte, 1024)
			message, errRead := r.Read(buf)
      if errRead != nil {
        if errRead != io.EOF {
          if verbose {
              log.Println("READ ERROR: ",errRead)
          }
        }
        break
      }
      _, errWrite := w.Write(buf[:message])
      w.Flush()
      if errWrite != nil {
        if verbose {
          log.Println("WRITE ERROR: ",errWrite)
        }
        return
      }
    }
  }()
  return c
}
