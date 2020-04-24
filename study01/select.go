package main

import (
	"fmt"
	"time"
)

/*

   每个 case 都必须是一个通信 （-<）
   所有 channel 表达式都会被求值(所有case都会被执行)
   所有被发送的表达式都会被求值
   如果任意某个通信可以进行，它就执行，其他被忽略。
   如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
   否则：
       如果有 default 子句，则执行该语句。
       如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

*/

func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
		fmt.Println(stopCh)
	}
	stopCh <- true
}

func main() {

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Recvice", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopCh:
			fmt.Println("Receive")
			//goto end
		case d := <-ch:
			fmt.Println("Receiveddd", d)
		}
	}
	//end:
}
