package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 定义一个整数类型
	var num1 int = 2300
	fmt.Println(num1)

	var num2 int8 = -20
	fmt.Println(num2)

	var num3 = 28
	// Printf 函数的作用就是: 格式化，把num3的类型 填充到 %T的位置上
	fmt.Printf("num3的类型是：%T ", num3) //num3 对应的类型
	fmt.Println()
	fmt.Println("占用字节数：", unsafe.Sizeof(num3))

	// 表示学生的年龄
	var age uint8 = 18
	fmt.Println(age)
}
