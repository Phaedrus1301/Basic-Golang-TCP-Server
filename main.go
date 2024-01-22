package main

import (
	"log"
	"net"
	"time"
)

func doThis(conn net.Conn) {
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	//do some processing here
	log.Println("processing stuff")
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World!\r\n"))
	conn.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("waiting for client to connect")

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")
		go doThis(conn)
	}
}
