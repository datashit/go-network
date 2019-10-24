package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", ":9090")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("connection established", conn.RemoteAddr().String())

	go listenTCP(conn)
	buf := []byte("hello tcp-server")
	for {
		time.Sleep(5 * time.Second)

		_, err := conn.Write(buf)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func listenTCP(conn net.Conn) {

	var buf [30]byte
	for {
		rlen, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(buf[:rlen]))
	}

}
