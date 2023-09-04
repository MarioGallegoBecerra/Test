package mksocket

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	SERVER_HOST = "192.168.0.107"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func Start() {
	fmt.Println("Server Running...")
	// start server
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close server when Start method end
	defer server.Close()

	// Listening for clients
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// If client connected
		go processClient(connection)
	}
}
func processClient(connection net.Conn) {
	defer connection.Close()
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

	response := "{\"hola\":\"holaa\"}"

	// Construir una respuesta HTTP válida
	httpResponse := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: application/json\r\n" +
		"Content-Length: " + strconv.Itoa(len(response)) + "\r\n" +
		"\r\n" +
		response

	_, err = connection.Write([]byte(httpResponse))
	if err != nil {
		fmt.Println("Error sending response: ", err.Error())
		return
	}

	fmt.Println("Sent")
}
