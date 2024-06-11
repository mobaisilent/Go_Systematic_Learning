package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	// 定义变量

	// for循环与select语句结合使用，使函数能够等待channel上的事件。
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// 两个int类型的chan

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		// 从chan c中接收并且打印6个数据
		quit <- 0
		// 向chan quit中发送1个数据（信号0）
	}()

	fibonacci(c, quit)
	// 因为上面的func 是一个 go 线程
}
