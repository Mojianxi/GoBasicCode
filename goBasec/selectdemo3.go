package main

//超时控制的经典实现
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("come to read ch!")
	case <-time.After(time.Second):
		fmt.Println("come to timeout")
	}
	fmt.Print("end of code!")
}
