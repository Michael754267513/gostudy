package main

import "fmt"

func main() {

	r := []int{1, 2, 3, 4, 5}
	for i, v := range r {
		fmt.Println(i, v)
	}
	for _, v := range r {
		print(v)
	}
}
