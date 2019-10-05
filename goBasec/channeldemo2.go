package main

import (
	"fmt"
	"time"
)

var ch chan int

func test_channel() {
	// ch := make(chan int)
	ch <- 1
	fmt.Println("come to end goroutline 1")
}
func main() {
	ch = make(chan int, 2) //值是1和值是0的适合输出执行顺序不同
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Println("running end!")
	<-ch
	time.Sleep(time.Second)
}
