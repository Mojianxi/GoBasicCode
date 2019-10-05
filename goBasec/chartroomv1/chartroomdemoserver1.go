package main

import (
	"fmt"
	"net"
	// "os"
)

func CheckError(err error) {
	if err != nil {
		// fmt.Println("Error :%s",err.Error())
		// os.Exit(1)
		panic(err)
	}
}
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		numofBytes, err := conn.Read(buf)
		// CheckError(err)
		if err != nil {
			// continue
			break
		}
		if numofBytes != 0 {
			fmt.Printf("Has received this message: %s", string(buf))
		}
	}
}
func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8062")
	CheckError(err)
	defer listen_socket.Close()
	fmt.Println("Server is wating...")
	for {
		conn, err := listen_socket.Accept()
		CheckError(err)
		//开一个协程处理
		go ProcessInfo(conn)
	}
}
