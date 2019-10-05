package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LOG_DIRECTORY = "./test.log"
)

var onlineConns = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)
var logFile *os.File
var logger *log.Logger

func CheckError(err error) {
	if err != nil {
		// fmt.Println("Error :%s",err.Error())
		// os.Exit(1)
		panic(err)
	}
}
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	// defer conn.Close()
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConns, addr)
		conn.Close()
		for i := range onlineConns {
			fmt.Println("closed" + i)
		}
	}(conn)
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
		// sendMessage := contents[1]
		sendMessage := strings.Join(contents[1:], "#")
		//防止有空格
		addr = strings.Trim(addr, " ")
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns sned failure!")
			}
		}
	} else {
		contents := strings.Split(message, "*")
		if strings.ToUpper(contents[1]) == "LIST" {
			var ips string = ""
			for i := range onlineConns {
				ips = ips + "|" + i
			}
			if conn, ok := onlineConns[contents[0]]; ok {
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("online conns send failure")
				}
			}
		}
	}
}
func main() {
	// onlineConns=make(map(string)net.Conn)
	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("logFile create failure")
		os.Exit(-1)
	}
	defer logFile.Close()
	logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8062")
	CheckError(err)
	defer listen_socket.Close()
	fmt.Println("Server is wating...")
	logger.Println("I am logger test")
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
