package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type people1 struct {
	// 结构体字段首字母必须大写
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Like []string `json:"like"`
}

func main() {
	p1 := []people1{{"曾浩", 20, []string{"apple", "game"}}, {"michael", 17, []string{"game", "orange"}}}
	json_data, err := json.Marshal(p1)
	if err != nil {
		panic("转换失败")
	} else {
		fmt.Println(string(json_data))
	}
	// 文件写入
	jfile, err := os.Create("./test.josn")
	if err != nil {
		panic("打开文件失败")
	}
	defer jfile.Close()

	enc_json := json.NewEncoder(jfile)
	err = enc_json.Encode(string(json_data))
	if err != nil {
		panic("文件存储错误")
	} else {
		fmt.Println("存储成功")
	}
	//// 读取文件
	//var p2 []people1
	//jf1 ,err := os.Open("C:\\Users\\hzeng\\go\\src\\gostudy.git\\test.josn")
	//if err != nil {
	//	panic("打开文件失败")
	//}
	//defer jf1.Close()
	//d_data := json.NewDecoder(jf1)
	//err1 := d_data.Decode(&p2)
	//if err1 != nil {
	//	fmt.Println("错误:",err1)
	//}else {
	//	fmt.Println(p2)
	//}
}
