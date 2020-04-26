package main

import "fmt"

func main() {
	/*
		 for true or flase {}
		 for init; condition; post { }
		for condition { }
		for { }
	*/
	a := 10
	for true {
		fmt.Println(a)
		break
	}
	for a <= 10 {
		fmt.Println(a)
		break // 终止循环
	}
	sum := 0
	for i := 0; i <= 10; i++ {
		if i == 9 {
			continue // 跳过此次循环
		}
		sum += i
	}
	// goto
	ischeck := 1
printname:
	fmt.Println("zenghao")
	for true {
		ischeck += 1
		if ischeck == 3 {
			break
		}
		goto printname
	}

	fmt.Println(sum)
	// 数组遍历
	name := []string{"michael", "zenghao"} // [number]<type>{value} 数组长度 类型 值
	for i, v := range name {
		fmt.Println(i, v)
	}

	number := [4]int{1, 2, 3}
	for i, v := range number {
		fmt.Println(i, v)
	}

	// for嵌套
	vara := true
	for true {
		for vara {
			fmt.Println(vara)
			break
		}
		vara := false
		fmt.Println(vara)
		break
	}
}
