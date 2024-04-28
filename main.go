package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "7000"
	connType = "tcp"
)

func main() {
	fmt.Println("Connecting to", connType, "server", connHost+":"+connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin:", err.Error())
			continue
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending data:", err.Error())
			continue
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err.Error())
			continue
		}

		log.Print("Server relay: " + message)
	}
}

//CID123456789,H123456789,C987654321,00:25:00,2024/01/22 13:45:00,2023/12/15 09:30:25,2022/06/30 16:59:59,Customer disconnected,0123456789,9876543210,100,200,1234567890,Reached voicemail,0987654321,300,Standard,0.05,1.50,Regular Call Plan,Direct,Mobile,Landline,VoIP,John Doe,Jane Smith,Generic VoIP,0
