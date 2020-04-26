package main

import (
	"fmt"
)

type people struct {
	name   string
	age    int
	sex    bool
	height float32
	weight float32
}

func main() {
	var p1 people

	p1.name = "曾浩"
	p1.age = 18
	p1.weight = 124.43
	fmt.Println(p1)
	fmt.Println(str(p1))
	strPtr(&p1)
}

// 结构体函数
func str(ren people) people {
	ren.height = 300
	return ren
}

// 结构体指针
func strPtr(ren *people) {
	fmt.Println(ren.name)
}
