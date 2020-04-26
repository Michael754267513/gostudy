package main

import "fmt"

/*
	需要设置退出条件，否则递归将陷入无限循环中
*/

func main() {
	fmt.Println(res(5))
}

func res(i int) int {
	ischeck := i
	ischeck++
	if ischeck == 10 {
		return ischeck
	}
	return res(ischeck)
}
