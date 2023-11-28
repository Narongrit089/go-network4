package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// Function send file to server
func sendFile(conn net.Conn, filePath string) {
	defer conn.Close() //Close connection before exit

	//Open file to send
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //Close file before exit

	//send file name to server
	_, fileName := filePath.Split(filePath)

	conn.Write([]byte(fileName)) //send file name to server

	//Create Buffer to read file
	buffer := make([]byte, 1024) //buffer size 1024 byte
	for {
		//Read file to buffer
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Send file success")
			} else {
				fmt.Println(err)
			}
			return
		}

		//Send Buffer to serer
		conn.Write(buffer[:n])
	}

}

func main() {
	//Connect to server
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close() //Close connection before exit
	//Print message connect success
	fmt.Println("Connect to server success")

	reader := bufio.NewReader(os.Stdin)
	//Get File Path and file Name from user
	fmt.Print("Enter file path+name: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	sendFile(conn, filePath)
}
