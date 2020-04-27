package main

import (
	"fmt"
	"time"
)

/*
	go 函数名( 参数列表 )

	runtime.GOMAXPROCS(逻辑CPU数量)  调整并发的运行性能
*/

func main() {
	//必须使用 make 创建 channel，
	//c1 := make(chan int)
	//fmt.Println(c1)
	go running()
	fmt.Println("test")

	// 匿名函数
	go func() {
		var t int
		for {
			t++
			fmt.Println(t)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func running() {
	var t int
	for {
		t++
		fmt.Println(t)
		time.Sleep(time.Second)
	}
}
