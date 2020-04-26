package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 0
	twoSum(a, b)
	fmt.Println("over")
}

func twoSum(num1 int, num2 int) int {
	// 捕捉异常
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	//手动抛出异常
	if num2 == 0 {
		panic(errors.New("分母不能为0"))
	}
	num1 = num1 / num2
	return num1
}
