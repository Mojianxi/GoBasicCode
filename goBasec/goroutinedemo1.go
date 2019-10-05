package main

import (
	"fmt"
	// "runtime"
	"strconv"
	"time"
)

func main() {
	//协程1
	go func() {
		for i := 1; i < 100; i++ {
			if i == 10 {
				// runtime.Gosched() //主动要求切换CPU
				<-ch //遇到阻塞主动让出,否则一直执行下去
			}
			fmt.Println("routine 1:" + strconv.Itoa(i))
		}
	}()
	//协程2
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("routine 2:" + strconv.Itoa(i))
		}
	}()
	time.Sleep(time.Second)
}
