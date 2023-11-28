package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024) //1024 bytes

	//Receive fileName from Client
	fileNameBuffer := make([]byte, 64) // 64 bytes

	n, err := conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := string(fileNameBuffer[:n])
	fmt.Println("Receive File Name:", fileName)

	//create file to store data
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //close file bofore exit

	//Receive and write data to file
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Transfer Complete")
			} else {
				fmt.Println(err)
			}
			return
		}
		//write data to file
		file.Write(buffer[:n])
	}

}

func main() {
	//Create Listener
	listener, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 5000")

	//Accept connection from Client
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Print Client address
		fmt.Println("Client Connected:", conn.RemoteAddr())
		//Handle connection
		go handleConnection(conn)
	}

}
