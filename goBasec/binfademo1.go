package main

import (
	"fmt"
	"strconv"
	"time"
)

func test_Routine() {
	fmt.Println("This is one routine!")
}
func Read(ch chan int) {
	value := <-ch
	fmt.Println("value:" + strconv.Itoa(value))
}
func Write(ch chan int) {
	// ch < -10
}
func main() {
	// go test_Routine()
	// time.Sleep(1)
	ch := make(chan int)
	//chan 的特性,阻塞 除非有goroutine对其进行操作
	go Read(ch)
	go Write(ch)
	fmt.Println("end of code")
	time.Sleep(10)
}
