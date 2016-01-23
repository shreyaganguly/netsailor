package main

import (
 "log"
 "net"
 "os"
 "io"

)

const (
  BUFFER_LIMIT = 2<<16 -1
)
func HandleUDPConnection(con *net.UDPConn) {


  chan_remote := readAndWriteUDP(con, os.Stdout,con, nil)
  ra := con.RemoteAddr()
  if ra == nil {
    ra = <- chan_remote
  }
  chan_local := readAndWriteUDP(os.Stdin,con,con,ra)
  select {
  case <- chan_local:
    log.Println("Connection closed from local process")
  case <- chan_remote:
      log.Println("Connection closed from remote process")
  }
}

func readAndWriteUDP(r io.Reader, w io.Writer, con *net.UDPConn, ra net.Addr) <-chan net.Addr {
	buf := make([]byte, BUFFER_LIMIT)
	cAddr := make(chan net.Addr)
	go func() {
		defer func() {
			con.Close()
			cAddr <- ra
		}()

		for {
			var bytesread int
			var errRead,errWrite error
			if con, ok := r.(*net.UDPConn); ok {
				var addr net.Addr
				bytesread, addr, errRead = con.ReadFrom(buf)
				if con.RemoteAddr() == nil && ra == nil {
					ra = addr
					cAddr <- ra
				}
			} else {
				bytesread, errRead = r.Read(buf)
			}
			if errRead != nil {
				if errRead != io.EOF {
					log.Println("READ ERROR: ",errRead)
				}
				break
			}
			if con, ok := w.(*net.UDPConn); ok && con.RemoteAddr() == nil {
				_, errWrite = con.WriteTo(buf[0:bytesread], ra)
			} else {
				_, errWrite = w.Write(buf[0:bytesread])
			}
			if errWrite != nil {
				log.Println("WRITE ERROR: ",errWrite)
        return
			}
		}
	}()
	return cAddr
}
