package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf) // this is a blocking call until the client sends data
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Processing the request")
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n")) // this is a blocking call until the client reads the data
	conn.Close()
}

func main() {
	listner, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for a connection...")
		conn, err := listner.Accept() // this is a blocking call until a client connects
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Got a connection")
		// multi thread server 
		// becoz we don't have to wait for the client to read the data 
		go do(conn)
	}
}
