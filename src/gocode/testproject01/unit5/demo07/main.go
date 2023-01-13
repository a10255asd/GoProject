package main

import "fmt"

// 参数的类型是指针 *代表地址
func test(num *int) {
	// 对地址对应的变量进行改变数值
	*num = 30
}

func main() {
	var num int = 10
	// 调用test函数的时候，传入num的地址 &
	fmt.Println("地址----", &num)
	test(&num)
	fmt.Println("main----", num)
}
