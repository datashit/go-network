package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp4", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	listentcp, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer listentcp.Close()

	fmt.Println("tcp server listen", addr.String())
	for {
		conn, err := listentcp.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}

		go handleTCP(conn)
	}
}

func handleTCP(conn *net.TCPConn) {

	send := []byte("hello tcp-client")
	var buf [30]byte
	for {

		rlen, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}

		_, err = conn.Write(send)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(buf[:rlen]))
	}
}
