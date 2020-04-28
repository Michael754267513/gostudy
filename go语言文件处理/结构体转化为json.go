package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	// 结构体字段首字母必须大写
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Like []string `json:"like"`
}

func main() {
	p1 := []people{{"曾浩", 20, []string{"apple", "game"}}, {"michael", 17, []string{"game", "orange"}}}
	fmt.Println(p1)
	json_data, err := json.Marshal(p1)
	if err != nil {
		panic("转换失败")
	} else {
		fmt.Println(string(json_data))
	}
}
