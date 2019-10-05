package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func MessageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client connect failure:" + err.Error())
			break
		}
	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8062")
	CheckError(err)
	defer conn.Close()
	//协程处理消息
	go MessageSend(conn)
	// conn.Write([]byte("Hello golang"))
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		CheckError(err)
		fmt.Println("receive server message connent:" + string(buf))
	}
	fmt.Println("client program end!")

}
