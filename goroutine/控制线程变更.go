package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	count := 1

	// 控制线程交替执行的变量
	flag := 0

	// 线程 1
	go func() {
		for count <= 100 {
			mu.Lock()
			for flag != 0 {
				cond.Wait()
			}
			if count > 100 {
				mu.Unlock()
				return
			}
			fmt.Println("线程 1:", count)
			count++
			flag = 1
			cond.Signal()
			mu.Unlock()
		}
	}()

	// 线程 2
	go func() {
		for count <= 100 {
			mu.Lock()
			for flag != 1 {
				cond.Wait()
			}
			if count > 100 {
				mu.Unlock()
				return
			}
			fmt.Println("线程 2:", count)
			count++
			flag = 0
			cond.Signal()
			mu.Unlock()
		}
	}()

	// 防止主 Goroutine 提前退出
	select {}
}
