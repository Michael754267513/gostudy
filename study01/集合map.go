package main

import (
	"fmt"
)

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值
Map 是使用 hash 表来实现
*/

func main() {
	peopleMap := make(map[string]string)
	peopleMap["name"] = "曾浩"
	peopleMap["age"] = "18"
	for _, v := range peopleMap {
		fmt.Println(v, peopleMap[v])
	}
	// 判断map中是否存在索引
	value, ischeck := peopleMap["name"]
	if ischeck {
		fmt.Println("存在", value)
	} else {
		fmt.Println("不存在", value)
	}
	// map 删除元素
	fmt.Println(peopleMap)
	delete(peopleMap, "age")
	fmt.Println(peopleMap)
}
