package main

//超时控制的经典实现
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- 1
	}()
	select {
	case <-ch:
		fmt.Println("come to read ch!")
	case <-timeout:
		fmt.Println("come to timeout")
	}
	fmt.Print("end of code!")
}
