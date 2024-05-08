`sync.Once`类型用于确保某个操作只执行一次，无论在并发环境下被多少个goroutine调用。它通常用于初始化操作或者只需执行一次的函数。

以下是`sync.Once`的基本使用示例：

```go
package main

import (
    "fmt"
    "sync"
)

var (
    once     sync.Once
    initialized bool
)

func initialize() {
    // 执行初始化操作
    fmt.Println("Initializing...")
    initialized = true
}

func main() {
    // 多个goroutine并发调用initialize函数，但只会执行一次
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(initialize)
        }()
    }
    // 等待goroutine执行完毕
    // 这里简单等待，实际应用中可能需要更加复杂的同步方式
    fmt.Scanln()

    // 检查是否已经初始化
    fmt.Println("Initialized:", initialized)
}
```

在上面的示例中，`sync.Once`类型的变量`once`用于确保`initialize`函数只被调用一次。多个goroutine并发调用`once.Do(initialize)`，但只有第一个调用会执行`initialize`函数，后续的调用会被忽略。