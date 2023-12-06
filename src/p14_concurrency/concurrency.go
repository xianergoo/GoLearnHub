// 进程示例
package main

//在操作系统中，进程可以理解为运行中的程序实例，
//线程是进程内的执行单元，而协程是轻量级的线程。
//它们之间的区别可以通过以下示例说明：

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("This is a process.")
    os.Exit(0) // 进程结束
}

// 线程示例
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("This is a thread.")
    }()
    wg.Wait() // 等待线程执行完成
}

// 协程示例
package main

import (
    "fmt"
    "time"
)

func main() {
    go printMessage()
    time.Sleep(time.Second) // 等待协程执行完成
}

func printMessage() {
    fmt.Println("This is a coroutine.")
}


// 并行示例
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 1; i <= 10; i++ {
            fmt.Println("Task A:", i)
        }
    }()
    go func() {
        defer wg.Done()
        for i := 1; i <= 10; i++ {
            fmt.Println("Task B:", i)
        }
    }()
    wg.Wait() // 等待两个任务执行完成
}

// 并发示例
package main

import (
    "fmt"
)

func main() {
    for i := 1; i <= 10; i++ {
        go fmt.Println("Task:", i)
    }
    // 主协程继续执行其他操作
}
