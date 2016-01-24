package main

import (
  "log"
  "net"
  "os"
  "bufio"
  "io"
)

func HandleTCPConnection(con *net.Conn,verbose bool) {
  chan_local := readAndWriteTCP(bufio.NewReader(os.Stdin), bufio.NewWriter(*con), con,verbose)
  chan_remote := readAndWriteTCP(bufio.NewReader(*con), bufio.NewWriter(os.Stdout), con,verbose)
  SelectChannel(chan_local,chan_remote,verbose)
}

func readAndWriteTCP(r *bufio.Reader, w *bufio.Writer, con *net.Conn, verbose bool) <-chan bool  {
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
