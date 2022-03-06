package main

import (
	"fmt"
	"log"
	"net"
	"path/filepath"

	"simple-ftp-server/ftp"
)

func main() {
	fmt.Println("start listen :8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs("files")
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
