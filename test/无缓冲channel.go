package main

import (
	"fmt"
	"time"
)

func NowChan() {
	c := make(chan int, 0) //无缓冲的通道

	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子协程结束")

		for i := 0; i < 3; i++ {
			fmt.Printf("子协程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
			c <- i //备注：如果在上面的话, 不会执行最后一次Printf
		}
	}()

	time.Sleep(2 * time.Second) //延时2s

	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main协程结束")
}

//有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。
//只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。

// 有缓冲的通道和无缓冲的通道之间的不同：
// 1）无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；
// 2）有缓冲的通道没有这种保证
func UseChan() {
	c := make(chan int, 3)
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))
	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子协程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("main协程结束")
}

//注意点：
//channel不像文件一样需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的，才去关闭channel；
//关闭channel后，无法向channel 再发送数据(引发 panic 错误后导致接收立即返回零值)；
//关闭channel后，可以继续向channel接收数据；
//  对于nil channel，无论收发都会被阻塞。

// * close()的使用
func ChanClose() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("Finished")
}

func main() {
	ChanClose()
}
