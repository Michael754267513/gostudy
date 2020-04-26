package main

import "fmt"

func main() {
	/*
		var variable_name [SIZE] variable_type
	*/
	s1 := [9]string{"a", "b"}
	s2 := []string{}
	fmt.Println(s1[1]) // 获取数组元素， 0开始
	s1[2] = "c"
	s1[4] = "d"
	s2 = append(s2, "ddd")
	//s1 = append(s1,"dddd")
	for i, v := range s2 {
		fmt.Println(i, v)
	}

	// 多维数组  var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type  x 为行，y 为列
	sz := [3][2]int{{1, 3}, {2, 10}, {1}}
	fmt.Println(sz)
	fmt.Println(sz[1][1]) // 二维数组访问
	// 多维数组
	sz1 := [3][3][2][2]int{}
	fmt.Println(sz1)
	fmt.Println(sz1[0])

	//函数传递数组
	arr := []int{1, 2, 3}
	fmt.Println(returnAvg(arr))
	fmt.Println(avg(arr))
}

func returnAvg(agv []int) []int {
	return agv
}

func avg(agv []int) int {
	sum := 0
	for i, v := range agv {
		fmt.Println(i)
		sum += v
	}
	return sum
}
