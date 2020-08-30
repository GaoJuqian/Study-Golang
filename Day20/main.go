package main

import (
	"fmt"
	"sync"
)

//	通道 channel

/*
	通道一计算平方
	通道二接收
*/
var wg = sync.WaitGroup{}

var test1 chan int
var test2 chan int

func main() {
	test1 = make(chan int, 1000)
	test2 = make(chan int, 2000)

	wg.Add(2)

	go go1()
	go go2(test1)

	wg.Wait() // 阻塞
	// 打印通道2平方的值
	for value := range test2 {
		fmt.Printf("test2的len:%v\n", len(test2))
		fmt.Println(value)
	}
}

// 生成 1～100放入通道test1
func go1() {
	for i := 1; i <= 1000; i++ {
		test1 <- i
	}
	close(test1) // 关闭通道1
	wg.Done()
}

// 计算通道1中的平方
func go2(v <-chan int) {
	for {
		value, ok := <-v // 从通道1中取值 如果为空 接收会阻塞
		if !ok {         // 无值false，break
			break
		}
		test2 <- value * value // 将平方存入通道test2
	}
	close(test2) // 关闭通道2
	wg.Done()

}
