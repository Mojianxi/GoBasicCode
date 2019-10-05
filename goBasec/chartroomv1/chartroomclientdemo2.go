package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8062")
	CheckError(err)
	defer conn.Close()
	conn.Write([]byte("Hello golang"))
	fmt.Println("Has sent the message!")
}
