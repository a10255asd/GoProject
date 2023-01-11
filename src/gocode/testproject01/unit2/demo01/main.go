package main

import "fmt"

func main() {
	// 变量声明
	var age int
	// 变量赋值
	age = 20
	// 变量使用
	fmt.Println("age=", age)
	// var age1 int = 20
	// fmt.Println(age1)
	// 声明和赋值可以合成一句
	var age2 int = 19
	fmt.Println("age2", age2)

	// var age int = 15
	// fmt.Println("age", age)
	// 变量的重复定义会报错

	// var num int = 12.56
	// fmt.Println(num)
	// 不可以在赋值的时候给予不匹配的类型
}
