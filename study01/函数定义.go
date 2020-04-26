package main

import (
	"fmt"
)

/*
func function_name( [parameter list] ) [return_types] {
   函数体
}
*/

func main() {
	var num1, num2 int = 10, 20
	fmt.Println(twoNumberSum(num1, num2))
	fmt.Println(returnThreeStr("曾", "浩"))
	swap(&num1, &num2) /* 保存 a 地址 */
	fmt.Println(num1, num2)
}

func twoNumberSum(num1, num2 int) int {

	var res int
	res = num1 + num2
	return res
}

// 多值返回

func returnThreeStr(s1, s2 string) (string, string, string) {
	return s1, s2, "哈罗"
}

// 引用传递， 实际是修改参数地址

func swap(a *int, b *int) {
	var temp int
	temp = *a
	fmt.Println(*a) /* 保存 a 地址上的值 */
	*a = *b
	*b = temp

}
