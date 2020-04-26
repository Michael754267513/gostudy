package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[1:2])
	fmt.Println(a[:2])
	fmt.Println(a[2:])
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	// 初始化空切片
	b := make([]int, 3, 3)
	b = append(b, 6) // 插入元素
	b = append(b, 6, 7, 8, 9)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	d := make([]int, 5, 5)
	copy(d, b) // 数组拷贝 注意长度和容量
	fmt.Println(d)
}
