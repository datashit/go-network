package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	listenudp, err := net.ListenUDP("udp4", addr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("udp server listen", addr.String())

	defer listenudp.Close()

	var buf [30]byte

	send := []byte("hello udp-client")
	for {
		rlen, raddr, err := listenudp.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatal(err)
		}

		_, err = listenudp.WriteToUDP(send, raddr)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(buf[:rlen]))

	}
}
