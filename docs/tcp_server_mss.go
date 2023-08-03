package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <port> <mss_value>\n", os.Args[0])
		os.Exit(1)
	}

	serverPort := os.Args[1]
	mssValue, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid MSS value")
		os.Exit(1)
	}

	address := fmt.Sprintf("0.0.0.0:%s", serverPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s with MSS value %d\n", serverPort, mssValue)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connection received from %s\n", conn.RemoteAddr().String())

	// Set MSS value
	tcpConn := conn.(*net.TCPConn)
	file, err := tcpConn.File()
	if err != nil {
		fmt.Println("Error getting socket file:", err)
		os.Exit(1)
	}
	defer file.Close()

	err = syscall.SetsockoptInt(int(file.Fd()), syscall.IPPROTO_TCP, syscall.TCP_MAXSEG, mssValue)
	if err != nil {
		fmt.Println("Error setting MSS value:", err)
		os.Exit(1)
	}

	welcomeMsg := []byte("Welcome")
	_, err = conn.Write(welcomeMsg)
	if err != nil {
		fmt.Println("Error sending welcome message:", err)
		os.Exit(1)
	}
}
