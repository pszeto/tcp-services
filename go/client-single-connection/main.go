package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func timer(interval int) {
	for {
		//check for time to be 0
		if interval <= 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
			//decrement interval every second
			interval--
		}
	}
}

func main() {
	interval, ok := os.LookupEnv("INTERVAL")
	if !ok {
		log.Println("INTERVAL not defined.  Defaulting to 300s")
		interval = "300"
	}
	intervalAsNum, err := strconv.Atoi(interval)
	if err != nil {
		log.Println("INTERVAL not a numnber.  Defaulting to 300")
		intervalAsNum = 300
	}
	serverAddress, ok := os.LookupEnv("UPSTREAM_ADDRESS")
	if !ok {
		log.Println("UPSTREAM_ADDRESS not defined.  Defaulting to tcp-server.tcp.svc.cluster.local:8000")
		serverAddress = "tcp-server.tcp.svc.cluster.local:8000"
	}

	podName, ok := os.LookupEnv("POD_NAME")
	podNamespace, ok := os.LookupEnv("POD_NAMESPACE")
	// Connect to the server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Println("Error connecting to: ", serverAddress)
		log.Println("Error:", err)
		return
	}

	// Enable keepalives
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		// keep alive
		tcpConn.SetKeepAlivePeriod(600 * time.Second)
	}

	for i := 0; i < 10000000; i++ {
		// Send some data to the server
		msg := "Hello, server! (iteration=" + strconv.Itoa(i) + ") " + podName + "." + podNamespace
		log.Println("Sending Message : ", msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Println("Error Sending Message:", err)
		}
		// Receive a response from the server
		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			log.Println("Error reading:", err)
		} else {
			log.Println("Server response:", string(reply))
		}
		timer(intervalAsNum)
	}
}
