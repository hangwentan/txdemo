
//好的，让我分别为您实现每种锁的示例代码：

//**互斥锁（Mutex）：**

//```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
//```

//**读写互斥锁（RWMutex）：**

//```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
//```

//**等待组（WaitGroup）：**

//```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    // 启动三个goroutine，并在每个goroutine中执行任务
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Println("Worker", id, "started")
            // 模拟任务执行
            fmt.Println("Worker", id, "finished")
        }(i)
    }

    // 等待所有goroutine执行完毕
    wg.Wait()
    fmt.Println("All workers have finished")
}
//```

//**条件变量（Cond）：**

//```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    var ready bool

    // 等待goroutine
    wg.Add(1)
    go func() {
        defer wg.Done()
        mu.Lock()
        for !ready {
            cond.Wait()
        }
        fmt.Println("Ready!")
        mu.Unlock()
    }()

    // 唤醒goroutine
    mu.Lock()
    ready = true
    cond.Signal()
    mu.Unlock()

    // 等待goroutine执行完毕
    wg.Wait()
}
```

好的，让我分别为您实现每种锁的示例代码：

**互斥锁（Mutex）：**

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
```

**读写互斥锁（RWMutex）：**

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
```

**等待组（WaitGroup）：**

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    // 启动三个goroutine，并在每个goroutine中执行任务
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Println("Worker", id, "started")
            // 模拟任务执行
            fmt.Println("Worker", id, "finished")
        }(i)
    }

    // 等待所有goroutine执行完毕
    wg.Wait()
    fmt.Println("All workers have finished")
}
```

**条件变量（Cond）：**

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    var ready bool

    // 等待goroutine
    wg.Add(1)
    go func() {
        defer wg.Done()
        mu.Lock()
        for !ready {
            cond.Wait()
        }
        fmt.Println("Ready!")
        mu.Unlock()
    }()

    // 唤醒goroutine
    mu.Lock()
    ready = true
    cond.Signal()
    mu.Unlock()

    // 等待goroutine执行完毕
    wg.Wait()
}
//```

// 这些示例代码演示了如何使用不同的锁和同步机制来管理并发访问和协作。这些示例代码演示了如何使用不同的锁和同步机制来管理并发访问和协作。