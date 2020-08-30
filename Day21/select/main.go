package main

import (
	"fmt"
	"time"
)

// select 满足哪个条件执行哪个通道
func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select { // 如果满足多个条件，随机执行一个
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
	time.Sleep(time.Second)
}
