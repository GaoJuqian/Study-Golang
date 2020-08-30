package main

import "fmt"

// work pool 线程池
func main() {
	// 使用线程池控制goroutine

	// 1. 初始化通道
	channel1 := make(chan int, 10) // 计算通道
	channel2 := make(chan int, 20) // 计算接收通道

	// 2. 开启goroutine
	for i := 1; i <= 3; i++ {
		go test(i, channel1, channel2)
	}

	// 3. 生成任务
	for i := 1; i <= 10; i++ {
		channel1 <- i
	}
	close(channel1) // 完毕，关闭通道1
	// 4. 打印chan2
	for i := 1; i <= 5; i++ {
		v := <-channel2
		fmt.Println(v)
	}
}

// work
func test(id int, chan1 <-chan int, chan2 chan<- int) {

	for v := range chan1 {
		fmt.Printf("start-id:%d,操作:%d\n", id, v)
		chan2 <- v * v
		fmt.Printf("end-id:%v\n", id)

	}
}
