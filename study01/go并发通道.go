package main

import (
	"fmt"
)

func main() {
	/*
		// 创建通道
		c1 := make(chan string)
		fmt.Println(c1)
		// 通道发送数据
		c1 <- "name" // 通道数据如果一直没有接受该通道会一直阻塞
		// 通道数据接受
		rev_c1 := <- c1 // 阻塞接受数据
		rev_c2, ok := <- c1 // 非祖泽接受数据
		<- c1 // 接受并扔掉数据
		fmt.Println(rev_c1)
		fmt.Println(rev_c2,ok)
	*/
	c1 := make(chan string)
	go func() {
		fmt.Println("start")
		c1 <- "name"
		fmt.Println("stop")
	}()
	fmt.Println("go func over")
	data, ok := <-c1
	fmt.Println(data, ok)
	// 循环接受通道数据
	c2 := make(chan int)
	times := 1
	go func() {
		for i := 0; i < 5; i++ {
			c2 <- times
			times++
		}
		c2 <- 0
	}()
	for data := range c2 {
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
}
