package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1
	}(ch)
	// time.Sleep(2)//加这一句执行结果不同
	select {
	case <-ch:
		fmt.Print("come to read ch!")
	default:
		fmt.Print("come to default!")
	}
}
