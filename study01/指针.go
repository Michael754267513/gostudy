package main

import "fmt"

/*
变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
*/

func main() {
	p := 10
	fmt.Println(p)
	fmt.Println("指针地址:", &p)
	// 指针申明 var var_name *var-type
	var ip *int
	a := 20
	ip = &a
	fmt.Println(ip)  //打印指针地址
	fmt.Println(*ip) //获取指针内存的值

	// 空指针
	var ptr *int
	fmt.Println(ptr)
	if ptr == nil {
		fmt.Println("空指针")
	}

	// 指针赋值
	num1 := []int{1, 2, 3, 5, 4}
	var ptr1 [3]*int
	for i := 0; i < 3; i++ {
		ptr1[i] = &num1[i]
	}
	for i, v := range ptr1 {
		fmt.Println(i, *v)
	}
}
