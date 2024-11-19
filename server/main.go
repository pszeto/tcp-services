package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Println("Port not defined.  Defaulting to 8000")
		port = "8000"
	}
	if _, err := strconv.Atoi(port); err != nil {
		log.Println("Port not a numnber.  Defaulting to 8000")
		port = "8000"
	}

	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":"+port)
	log.Println("TCP Server Started.  Listening on port:", port)
	// Close the listener when the application closes.
	defer ln.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	podName, _ := os.LookupEnv("POD_NAME")
	podNamespace, _ := os.LookupEnv("POD_NAMESPACE")
	for {
		// Read incoming data
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		// Print the incoming data
		log.Printf("Received: %s", buf)
		msg := "Message received! " + podName + "." + podNamespace
		conn.Write([]byte(msg))
	}
}
