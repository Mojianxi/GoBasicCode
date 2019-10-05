package main

import (
	"fmt"
	"net"
	"strings"
	// "os"
)

var onlineConns = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)

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
			// remoteAddr := conn.RemoteAddr()
			// fmt.Print(remoteAddr)
			// fmt.Printf("Has received this message: %s\n", string(buf[0:numofBytes]))
			message := string(buf[0:numofBytes])
			//每次收到消息,写入到messageQueue
			messageQueue <- message
		}
	}
}
func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break

		}
	}
}
func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		sendMessage := contents[1]
		//防止有空格
		addr = strings.Trim(addr, " ")
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns sned failure!")
			}
		}
	}
}
func main() {
	// onlineConns=make(map(string)net.Conn)

	listen_socket, err := net.Listen("tcp", "127.0.0.1:8062")
	CheckError(err)
	defer listen_socket.Close()
	fmt.Println("Server is wating...")
	//开启一个协程处理messageQueue
	go ConsumeMessage()
	for {
		conn, err := listen_socket.Accept()
		CheckError(err)
		//连接之前将conn对象存储到onlineConns映射表中
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn
		for i := range onlineConns {
			fmt.Println(i)
		}
		//开一个协程处理
		go ProcessInfo(conn)
	}
}
