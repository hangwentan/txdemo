package main

import (
    "fmt"
    "sync"
)

// 互斥锁
func mutex() {
    var mutex sync.Mutex
    var sharedResource int

    // goroutine 1
    go func() {
        mutex.Lock()
        sharedResource = 42
        mutex.Unlock()
    }()

    // goroutine 2
    go func() {
        mutex.Lock()
        fmt.Println("Shared resource:", sharedResource)
        mutex.Unlock()
    }()

    // 等待goroutine执行完毕
    fmt.Scanln()
}