package main

// 并发锁 互斥锁 读写锁
// 	sync.Once{} .Do() 加标志位保证并发时执行一次
import (
	"fmt"
	"sync"
	"time"
)

var (
	t      int
	wg     sync.WaitGroup
	lock   sync.Mutex   // 互斥锁
	rwlock sync.RWMutex // 读写互斥锁
)

func main() {
	wg.Add(3) // 等待 goroutine个数
	start := time.Now()
	go read()
	go add()
	go add()
	wg.Wait() // 阻塞主线程
	fmt.Println(t)
	fmt.Printf("时间：%v\n", time.Since(start))
}

func add() { // 写操作
	for i := 1; i <= 100; i++ {
		rwlock.Lock() // 读写锁 写锁
		//lock.Lock() // 互斥锁
		t += 1
		//lock.Unlock() // 互斥锁 解锁
		rwlock.Unlock() // 读写锁 写锁

	}
	wg.Done() // 结束当前goroutine
}
func read() { // 读操作
	rwlock.RLock() // 读锁
	//lock.Lock() // 互斥锁
	for i := 1; i <= 1000; i++ {
		fmt.Println(t)
	}
	//lock.Unlock() // 互斥锁 解锁
	rwlock.RUnlock() // 解锁

	wg.Done() // 结束当前goroutine

}
