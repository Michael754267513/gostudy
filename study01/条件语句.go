package main

import "fmt"

func main() {
	var a int = 10

	if a >= 5 {
		fmt.Println("大于5")
	}

	if a < 5 {
		fmt.Println("不存在")
	} else {
		fmt.Println("不小于5")
	}

	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	// fallthrough 强制执行后面的case   break终止
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
		fallthrough
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
