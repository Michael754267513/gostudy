package main

import "fmt"

/*
   定义一个接口
*/

func main() {
	var p1 sex
	p1 = new(man)
	fmt.Println(p1.age())
	fmt.Println(p1.weight(182.87))
	fmt.Println(new(woman).weight(222))
	p1.call()
	new(woman).call()
}

type sex interface {
	age() int
	weight(weight float32) float32
	call()
}

type man struct {
	name string
}

type woman struct {
	name string
	age  int
	face string
}

func (man man) call() {
	man.name = "男人"
	fmt.Println(man)
}

func (man man) weight(weight float32) float32 {
	return weight
}
func (woman woman) weight(weight float32) float32 {
	return weight
}

func (man man) age() int {
	return 18
}

func (woman woman) call() {
	woman.name = "女人"
	woman.face = "巨丑"
	fmt.Println(woman)
}
