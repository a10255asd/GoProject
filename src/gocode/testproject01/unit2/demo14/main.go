package main

import "fmt"

func main() {
	var age int = 18
	// & 符号 + 变量，就可以获取这个变量内存的地址
	fmt.Println(&age)
	// 定义一个指针变量
	// var 代表声明一个变量，ptr 指针变量的名字
	// ptr对应的类型是*int，是一个指针类型（指向int类型的指针）
	// &age就是一个地址，是ptr变量的具体的值
	var ptr *int = &age
	fmt.Println("地址是：", ptr)
	fmt.Println("ptr本身存储空间的地址", &ptr)
	// 获取ptr这个指针或者这个地址只想的那个数据
	fmt.Println("ptr指向的数值为:", *ptr)
}
