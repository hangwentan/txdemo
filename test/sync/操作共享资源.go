package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once   sync.Once
	status bool
)

func initialized() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("事件回退数据")
		}
	}()
	fmt.Println("修改数据")
	status = true
	panic("修改失败")
}

func SyncOnce() {

	ch := make(chan int, 0)

	var wait sync.WaitGroup

	for i := 0; i < 10; i++ {
		wait.Add(1)
		fmt.Println("处理次数:", i+1)
		go func() {
			ch <- i + 1
			once.Do(initialized)
			wait.Done()
		}()
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num:", num)
	}

	wait.Wait()

	fmt.Scanln()

	// 检查是否已经初始化
	fmt.Println("Initialized:", status)
}

func SyncRWMutex() {
	var rwMutex sync.RWMutex
	var sharedResource int

	// 读取共享资源的goroutine
	go func() {
		rwMutex.RLock()
		fmt.Println("Shared resource:", sharedResource)
		rwMutex.RUnlock()
	}()

	// 写入共享资源的goroutine
	go func() {
		rwMutex.Lock()
		sharedResource = 42
		rwMutex.Unlock()
	}()

	// 等待goroutine执行完毕
	fmt.Scanln()
}

func main() {
	//SyncOnce()
	SyncRWMutex()
}
