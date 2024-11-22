package main

import "fmt"

func main() {
	// 声明
	var pa, pb *int
	// 取址
	var a = 10 // 默认为 int
	pa = &a
	// 解引用（取值）
	var b = *pa
	// 输出
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("&a:", &a)
	fmt.Println("&b:", &b)
	fmt.Println("pa:", pa)
	fmt.Println("pb:", pb)
	fmt.Println("*pa:", *pa)
	// fmt.Println("*pb:", *pb)
}
