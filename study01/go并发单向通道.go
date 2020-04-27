package main

import "fmt"

func main() {
	ch := make(chan int)
	// 单项接受数据
	var revData <-chan int = ch
	// 单项发送数据
	var sendData chan<- int = ch
	fmt.Println(revData, sendData)
}
