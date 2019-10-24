package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp4", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("udp-client ready")
	go listenUDP(conn)

	buf := []byte("hello udp-server")
	for {
		time.Sleep(5 * time.Second)

		_, err := conn.Write(buf)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func listenUDP(conn net.Conn) {
	var buf [30]byte
	for {
		rlen, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(buf[:rlen]))
	}
}
