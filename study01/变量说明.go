package main

import "fmt"

/*
	1. bool 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
	2. 数字类型 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
    3. 字符串类型 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
	4. 派生类 指针类型（Pointer） 数组类型  结构化类型(struct) Channel 类型  函数类型 切片类型 接口类型（interface）  Map 类型
*/

// 数字类型  uint* int*    uint不包括负数 2^* (uint8 无符号 8 位整型 (0 到 255))   int包括负数  -2^*-1 ~ 2^*-1 (int8 有符号 8 位整型 (-128 到 127) )

// 浮点类型 float*   complex64 32 位实数和虚数  complex128 64 位实数和虚数

// other类型 byte->uint8   runne->int32 uint->32或者64位 int->uint  uintprt 无符号整型 存放指针

func main() {
	// 变量申明1 var identifier type
	//变量申明2 var identifier1, identifier2 type
	var a1, a2 string = "a1", "a2"
	var a int = 100
	fmt.Println(a)
	fmt.Println(a1 + " " + a2)
	// 系统默认值 bool为flase int为0 字符串为“”
	var c bool
	var d int
	var e string
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	/* 一下类型为nil
	var a *int
	var a []int
	var a map[string] int
	var a chan int
	var a func(string) int
	var a error // error 是接口
	*/
	var dd map[string]int
	fmt.Println(dd)
	fmt.Printf("%v %v %v %v %v %v %v\n", a, a1, a2, c, d, e, dd)
	// 未申明的变量无法单独使用 需要配合已经声明的变量进行操作 ：=
	e, name := "ee", "zenghao"
	fmt.Println(name)

	var i int = 10
	var j int = i
	fmt.Println(j)
	i = 100
	fmt.Println(j)

	// 常量使用方式 iota递增 << 为<<n==*(2^n)  左值不变，右为(2^n)
	const (
		aa = iota
		bb
		cc
	)
	fmt.Println(aa, bb, cc)
}
