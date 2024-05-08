// 在Go语言中，常见的锁包括以下几种：

// 1. **互斥锁（Mutex）**：`sync.Mutex`是Go语言标准库中最基本的锁类型之一。它提供了对共享资源的排他性访问，只有持有锁的goroutine可以访问被保护的资源，其他goroutine需要等待锁释放才能访问。

// 2. **读写互斥锁（RWMutex）**：`sync.RWMutex`是互斥锁的一种变种，它允许多个goroutine同时读取共享资源，但只有一个goroutine可以写入共享资源。在读多写少的场景中，使用读写互斥锁可以提高并发性能。

// 3. **等待组（WaitGroup）**：`sync.WaitGroup`用于等待一组goroutine完成执行。它提供了`Add`、`Done`和`Wait`方法，用于添加goroutine、标记goroutine完成和等待所有goroutine完成。

// 4. **条件变量（Cond）**：`sync.Cond`是一个条件变量，它可以让一个或多个goroutine等待某个条件的发生。条件变量通常与互斥锁配合使用，当条件不满足时，goroutine会进入等待状态，当条件满足时，通过条件变量唤醒等待的goroutine。

// 5. **原子操作（Atomic）**：`sync/atomic`包提供了一些原子操作函数，用于对内存地址进行原子操作，如增加、减少、交换等。原子操作是不可分割的操作，可以用于实现一些简单的锁机制和同步机制。

这些锁和同步机制在不同的场景下有不同的用途，开发者可以根据具体需求选择合适的锁来确保并发安全性。

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