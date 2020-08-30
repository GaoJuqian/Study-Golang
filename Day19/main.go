package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// 并发 goroutine
func main() {
	wg.Add(10000) // goroutines数量
	for i := 0; i < 10000; i++ {
		go func(i int) { // 由于闭包原因，线程调用时，循环已结束,(传参形式可使打印时的i固定)
			fmt.Println(i)
			wg.Done() // 单个goroutine结束
		}(i)
	}
	wg.Wait() // 阻塞主线程，直至 goroutines完毕

	fmt.Println("main out")

}
